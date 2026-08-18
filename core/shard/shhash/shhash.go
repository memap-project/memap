package shhash

import (
	"hash/fnv"

	"github.com/memap-project/memap/core/shard"
	"github.com/memap-project/memap/core/vals"
)

type ShardedHash struct {
	shardCount uint8
	shards     []*shard.Shard[*vals.Hash]
}

func NewShardedHash() *ShardedHash {
	shards := make([]*shard.Shard[*vals.Hash], 8)
	for i := range shards {
		shards[i] = shard.NewShard[*vals.Hash]()
	}
	return &ShardedHash{
		shardCount: 8,
		shards:     shards,
	}
}

// getShard returns the shard that the key belongs to.
func (s *ShardedHash) getShard(key string) *shard.Shard[*vals.Hash] {
	h := fnv.New32a()
	h.Write([]byte(key))
	idx := h.Sum32() & uint32(s.shardCount-1)
	return s.shards[idx]
}

// Get retrieves the hash for the given key.
// Returns an empty map and false if the hash does not exist.
func (s *ShardedHash) Get(key string) (map[string]string, bool) {
	h, ok := s.getShard(key).Get(key)
	if !ok {
		return map[string]string{}, false
	}
	if h.IsExpired() {
		s.getShard(key).Delete(key)
		return map[string]string{}, false
	}
	return h.GetCopy(), true
}

// Set creates or overwrites hash.
// Creates if the hash does not exist, or overwrites if it does.
func (s *ShardedHash) Set(key string, ttl int64) bool {
	h := vals.NewHash()
	if ttl > 0 {
		h.Expire(ttl)
	}
	s.getShard(key).Set(key, h)
	return true
}

// Delete deletes the hash map for the given key.
func (s *ShardedHash) Delete(key string) {
	s.getShard(key).Delete(key)
}

// Expire expires the hash for the given key.
func (s *ShardedHash) Expire(key string, ttl int64) bool {
	return s.getShard(key).Update(key, func(h *vals.Hash) bool {
		if h.IsExpired() {
			return false
		}
		h.Expire(ttl)
		return true
	})
}

// TTL returns the remaining time of the hash for the given key.
// Returns -1 and false if the hash has no expiration time.
// Returns -2 and false if the hash does not exist.
func (s *ShardedHash) TTL(key string) int64 {
	h, ok := s.getShard(key).Get(key)
	if !ok || h.IsExpired() {
		return -2
	}
	if h.TTL() == 0 {
		return -1
	}
	return h.TTL()
}

// Exists returns true if the hash for the given key exists.
func (s *ShardedHash) Exists(key string) bool {
	_, ok := s.getShard(key).Get(key)
	return ok
}

// LEN returns the number of fields in the hash for the given key.
// Returns 0 and false if the key does not exist.
func (s *ShardedHash) Len(key string) (int64, bool) {
	h, ok := s.getShard(key).Get(key)
	if !ok {
		return 0, false
	}
	if h.IsExpired() {
		s.getShard(key).Delete(key)
		return 0, false
	}
	return h.Len(), true
}

// Keys returns all the fields of the hash for the given key.
// Returns empty slice and false if the key does not exist.
func (s *ShardedHash) Keys(key string) ([]string, bool) {
	h, ok := s.getShard(key).Get(key)
	if !ok {
		return []string{}, false
	}
	if h.IsExpired() {
		s.getShard(key).Delete(key)
		return []string{}, false
	}
	return h.Keys(), true
}

// Values returns all the values of the hash for the given key.
// Returns empty slice and false if the key does not exist.
func (s *ShardedHash) Values(key string) ([]string, bool) {
	h, ok := s.getShard(key).Get(key)
	if !ok {
		return []string{}, false
	}
	if h.IsExpired() {
		s.getShard(key).Delete(key)
		return []string{}, false
	}
	return h.Values(), true
}

// GetField retrieves a value of the field from the hash for the given key and field.
// Returns empty string and false if the key or field does not exist.
func (s *ShardedHash) GetField(key string, field string) (string, bool) {
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

// SetField sets a field in the hash for the given key and field.
// Creates a new hash if not exists.
func (s *ShardedHash) SetField(key, field, value string) bool {
	h, _ := s.getShard(key).GetOrInit(key, vals.NewHash)
	if h.IsExpired() {
		h = vals.NewHash()
		s.getShard(key).Set(key, h)
	}
	h.Set(field, value)
	return true
}

// DeleteField deletes a field from the hash for the given key and field.
func (s *ShardedHash) DeleteField(key, field string) {
	h, ok := s.getShard(key).Get(key)
	if ok {
		h.Delete(field)
	}
}

// CleanExpired cleans expired hashes.
func (s *ShardedHash) CleanExpired() {
	for _, shard := range s.shards {
		shard.Clean(func(_ string, h *vals.Hash) bool {
			return h.IsExpired()
		})
	}
}

// Flush removes all hashes from the shard.
func (s *ShardedHash) Flush() {
	for _, shard := range s.shards {
		shard.Flush()
	}
}
