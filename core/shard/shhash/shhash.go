package shhash

import (
	"hash/fnv"

	"github.com/dmi3midd/memap/core/shard"
)

type ShardedHash struct {
	shardCount uint8
	shards     []*shard.Shard[*Hash]
}

func NewShardedHash() *ShardedHash {
	shards := make([]*shard.Shard[*Hash], 8)
	for i := range shards {
		shards[i] = shard.NewShard[*Hash]()
	}
	return &ShardedHash{
		shardCount: 8,
		shards:     shards,
	}
}

// getShard returns the shard that the key belongs to.
func (s *ShardedHash) getShard(key string) *shard.Shard[*Hash] {
	h := fnv.New32a()
	h.Write([]byte(key))
	idx := h.Sum32() & uint32(s.shardCount-1)
	return s.shards[idx]
}

// HGet retrieves the hash for the given key.
// Returns nil and false if the hash does not exist.
func (s *ShardedHash) HGet(key string) (map[string]string, bool) {
	h, ok := s.getShard(key).Get(key)
	if !ok {
		return nil, false
	}
	if h.IsExpired() {
		s.getShard(key).Delete(key)
		return nil, false
	}
	return h.GetCopy(), true
}

// HSet creates empty hash map if not exists.
func (s *ShardedHash) HSet(key string) {
	s.getShard(key).GetOrInit(key, NewHash)
}

// HDelete deletes the hash map for the given key.
func (s *ShardedHash) HDelete(key string) {
	s.getShard(key).Delete(key)
}

// HFGet retrieves a value of the field from the hash for the given key and field.
// Returns empty string and false if the key or field does not exist.
func (s *ShardedHash) HFGet(key string, field string) (string, bool) {
	h, ok := s.getShard(key).Get(key)
	if !ok {
		return "", false
	}
	if h.IsExpired() {
		s.getShard(key).Delete(key)
		return "", false
	}
	return h.Get(field)
}

// HFSet sets a field in the hash for the given key and field. Creates hash if not exists.
func (s *ShardedHash) HFSet(key, field, value string) {
	h := s.getShard(key).GetOrInit(key, NewHash)
	h.Set(field, value)
}

// HFDelete deletes a field from the hash for the given key and field.
func (s *ShardedHash) HFDelete(key, field string) {
	h, ok := s.getShard(key).Get(key)
	if ok {
		h.Delete(field)
	}
}

// Clean cleans expired hashes.
func (s *ShardedHash) Clean() {
	for _, shard := range s.shards {
		shard.Clean(func(_ string, h *Hash) bool {
			return h.IsExpired()
		})
	}
}
