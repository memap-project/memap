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

// HGet retrieves the hash for the given key.
// Returns an empty map and false if the hash does not exist.
func (s *ShardedHash) HGet(key string) (map[string]string, bool) {
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

// HSet creates or overwrites hash.
// Creates if the hash does not exist, or overwrites if it does.
func (s *ShardedHash) HSet(key string, ttl int64) bool {
	h := vals.NewHash()
	if ttl > 0 {
		h.Expire(ttl)
	}
	s.getShard(key).Set(key, h)
	return true
}

// HDelete deletes the hash map for the given key.
func (s *ShardedHash) HDelete(key string) {
	s.getShard(key).Delete(key)
}

// HExpire expires the hash for the given key.
func (s *ShardedHash) HExpire(key string, ttl int64) bool {
	return s.getShard(key).Update(key, func(h *vals.Hash) bool {
		if h.IsExpired() {
			return false
		}
		h.Expire(ttl)
		return true
	})
}

// HTTL returns the remaining time of the hash for the given key.
// Returns -1 and false if the hash has no expiration time.
// Returns -2 and false if the hash does not exist.
func (s *ShardedHash) HTTL(key string) int64 {
	h, ok := s.getShard(key).Get(key)
	if !ok || h.IsExpired() {
		return -2
	}
	if h.LeftTime() == 0 {
		return -1
	}
	return h.LeftTime()
}

// HExists returns true if the hash for the given key exists.
func (s *ShardedHash) HExists(key string) bool {
	_, ok := s.getShard(key).Get(key)
	return ok
}

// HLEN returns the number of fields in the hash for the given key.
// Returns 0 and false if the key does not exist.
func (s *ShardedHash) HLen(key string) (int64, bool) {
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

// HKeys returns all the fields of the hash for the given key.
// Returns empty slice and false if the key does not exist.
func (s *ShardedHash) HKeys(key string) ([]string, bool) {
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

// HValues returns all the values of the hash for the given key.
// Returns empty slice and false if the key does not exist.
func (s *ShardedHash) HValues(key string) ([]string, bool) {
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

// HFSet sets a field in the hash for the given key and field.
// Creates a new hash if not exists.
func (s *ShardedHash) HFSet(key, field, value string) bool {
	h, _ := s.getShard(key).GetOrInit(key, vals.NewHash)
	if h.IsExpired() {
		h = vals.NewHash()
		s.getShard(key).Set(key, h)
	}
	h.Set(field, value)
	return true
}

// HFDelete deletes a field from the hash for the given key and field.
func (s *ShardedHash) HFDelete(key, field string) {
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
