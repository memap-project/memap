package shmap

import (
	"hash/fnv"
	"sync"

	"github.com/dmi3midd/memap/core/item"
	"github.com/dmi3midd/memap/core/shard"
)

// ShardedMap is a map that stores shards of data for fast sharded access.
type ShardedMap struct {
	mu         sync.RWMutex
	shardCount int
	shards     []*shard.Shard
}

// NewShardedMap creates a new ShardedMap with shardCount shards.
func NewShardedMap() *ShardedMap {
	shards := make([]*shard.Shard, 8)
	for i := range shards {
		shards[i] = shard.NewShard()
	}
	return &ShardedMap{
		shardCount: 8,
		shards:     shards,
	}
}

// getShard returns the shard that the key belongs to.
func (s *ShardedMap) getShard(key string) *shard.Shard {
	h := fnv.New32a()
	h.Write([]byte(key))
	idx := h.Sum32() & uint32(s.shardCount-1)
	return s.shards[idx]
}

// Get retrieves an item from the sharded map.
func (s *ShardedMap) Get(key string) (*item.Item, bool) {
	item, ok := s.getShard(key).Get(key)
	return item, ok
}

// Set sets or updates an item in the sharded map.
func (s *ShardedMap) Set(key string, value item.Item) {
	s.getShard(key).Set(key, value)
}

// Delete deletes an item from the sharded map.
func (s *ShardedMap) Delete(key string) {
	s.getShard(key).Delete(key)
}
