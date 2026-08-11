package shmap

import (
	"hash/fnv"

	"github.com/memap-project/memap/core/item"
	"github.com/memap-project/memap/core/shard"
)

// ShardedMap is a map that stores shards of data for fast sharded access.
type ShardedMap struct {
	shardCount uint8
	shards     []*shard.Shard[item.Item]
}

// NewShardedMap creates a new ShardedMap with default shard count (8).
func NewShardedMap() *ShardedMap {
	shards := make([]*shard.Shard[item.Item], 8)
	for i := range shards {
		shards[i] = shard.NewShard[item.Item]()
	}
	return &ShardedMap{
		shardCount: 8,
		shards:     shards,
	}
}

// getShard returns the shard that the key belongs to.
func (s *ShardedMap) getShard(key string) *shard.Shard[item.Item] {
	h := fnv.New32a()
	h.Write([]byte(key))
	idx := h.Sum32() & uint32(s.shardCount-1)
	return s.shards[idx]
}

// Expire sets expiration time for the key if it exists and is not expired.
func (s *ShardedMap) Expire(key string, ttl int64) bool {
	return s.getShard(key).Update(key, func(it *item.Item) bool {
		if it.IsExpired() {
			return false
		}
		it.Expire(ttl)
		return true
	})
}

// Get retrieves an item from the sharded map.
// Returns zero value and false if the key doesn't exist.
func (s *ShardedMap) Get(key string) (item.Item, bool) {
	i, ok := s.getShard(key).Get(key)
	if !ok {
		return item.Item{}, false
	}
	if i.IsExpired() {
		s.getShard(key).Delete(key)
		return item.Item{}, false
	}
	return i, true
}

// Set sets or updates an item in the sharded map.
func (s *ShardedMap) Set(key string, value item.Item) {
	s.getShard(key).Set(key, value)
}

// Delete deletes an item from the sharded map.
func (s *ShardedMap) Delete(key string) {
	s.getShard(key).Delete(key)
}

// Clean cleans expired items.
func (s *ShardedMap) Clean() {
	for _, shard := range s.shards {
		shard.Clean(func(_ string, item item.Item) bool {
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
