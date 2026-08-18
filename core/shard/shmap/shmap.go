package shmap

import (
	"hash/fnv"

	"github.com/memap-project/memap/core/shard"
	"github.com/memap-project/memap/core/vals"
)

// ShardedMap is a map that stores shards of data for fast sharded access.
type ShardedMap struct {
	shardCount uint8
	shards     []*shard.Shard[*vals.Item]
}

// NewShardedMap creates a new ShardedMap with default shard count (8).
func NewShardedMap() *ShardedMap {
	shards := make([]*shard.Shard[*vals.Item], 8)
	for i := range shards {
		shards[i] = shard.NewShard[*vals.Item]()
	}
	return &ShardedMap{
		shardCount: 8,
		shards:     shards,
	}
}

// getShard returns the shard that the key belongs to.
func (s *ShardedMap) getShard(key string) *shard.Shard[*vals.Item] {
	h := fnv.New32a()
	h.Write([]byte(key))
	idx := h.Sum32() & uint32(s.shardCount-1)
	return s.shards[idx]
}

// Get retrieves an item from the sharded map.
// Returns empty string and false if the key doesn't exist.
func (s *ShardedMap) Get(key string) (string, bool) {
	shard := s.getShard(key)
	i, ok := shard.Get(key)
	if !ok {
		return "", false
	}
	if i.IsExpired() {
		shard.Delete(key)
		return "", false
	}
	return i.GetValue(), true
}

// Set upserts an item in the sharded map.
func (s *ShardedMap) Set(key, value string, ttl int64) bool {
	i := vals.NewItem()
	i.SetValue(value)
	if ttl > 0 {
		i.Expire(ttl)
	}
	s.getShard(key).Set(key, i)
	return true
}

// Delete deletes an item from the sharded map.
func (s *ShardedMap) Delete(key string) {
	s.getShard(key).Delete(key)
}

// Expire sets expiration time for the key if it exists.
func (s *ShardedMap) Expire(key string, ttl int64) bool {
	return s.getShard(key).Update(key, func(i *vals.Item) bool {
		if i.IsExpired() {
			return false
		}
		i.Expire(ttl)
		return true
	})
}

// TTL returns the time-to-live of the key if it exists and is not expired.
// Returns -1 and false if the key has no expiration time.
// Returns -2 and false if the key does not exist.
func (s *ShardedMap) TTL(key string) int64 {
	i, ok := s.getShard(key).Get(key)
	if !ok || i.IsExpired() {
		return -2
	}
	if i.LeftTime() == 0 {
		return -1
	}
	return i.LeftTime()
}

// CleanExpired cleans expired items.
func (s *ShardedMap) CleanExpired() {
	for _, shard := range s.shards {
		shard.Clean(func(key string, item *vals.Item) bool {
			return item.IsExpired()
		})
	}
}

// Flush removes all items from the shard.
func (s *ShardedMap) Flush() {
	for _, shard := range s.shards {
		shard.Flush()
	}
}
