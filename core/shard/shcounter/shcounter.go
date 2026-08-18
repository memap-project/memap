package shcounter

import (
	"hash/fnv"

	"github.com/memap-project/memap/core/shard"
	"github.com/memap-project/memap/core/vals"
)

type ShardedCounter struct {
	shardCount uint8
	shards     []*shard.Shard[*vals.Counter]
}

func NewShardedCounter() *ShardedCounter {
	shards := make([]*shard.Shard[*vals.Counter], 8)
	for i := range shards {
		shards[i] = shard.NewShard[*vals.Counter]()
	}
	return &ShardedCounter{
		shardCount: 8,
		shards:     shards,
	}
}

// getShard returns the shard that the key belongs to.
func (s *ShardedCounter) getShard(key string) *shard.Shard[*vals.Counter] {
	h := fnv.New32a()
	h.Write([]byte(key))
	idx := h.Sum32() & uint32(s.shardCount-1)
	return s.shards[idx]
}

// Init initializes a counter with the given key, limit, and ttl.
// Returns true if the counter was initialized, false if it already existed.
func (s *ShardedCounter) Init(key string, limit int64, ttl int64) bool {
	c, ok := s.getShard(key).GetOrInit(key, vals.NewCounter)
	if ok {
		return false
	}
	c.Expire(ttl)
	c.SetLimit(limit)
	return true
}

// SetLimit sets the limit of the counter with the given key.
// Returns true if the limit was set, false if the counter does not exist.
func (s *ShardedCounter) SetLimit(key string, limit int64) bool {
	shard := s.getShard(key)
	c, ok := shard.Get(key)
	if !ok {
		return false
	}
	if c.IsExpired() {
		shard.Delete(key)
		return false
	}
	c.SetLimit(limit)
	return true
}

// GetLimit returns the limit of the counter with the given key.
// Returns 0 and false if the counter does not exist.
func (s *ShardedCounter) GetLimit(key string) (int64, bool) {
	shard := s.getShard(key)
	c, ok := shard.Get(key)
	if !ok {
		return 0, false
	}
	if c.IsExpired() {
		shard.Delete(key)
		return 0, false
	}
	return c.GetLimit(), true
}

// Get returns the value of the counter with the given key.
// Returns 0 and false if the counter does not exist.
func (s *ShardedCounter) Get(key string) (int64, bool) {
	shard := s.getShard(key)
	c, ok := shard.Get(key)
	if !ok {
		return 0, false
	}
	if c.IsExpired() {
		shard.Delete(key)
		return 0, false
	}
	return c.GetValue(), true
}

// Delete deletes the counter with the given key.
func (s *ShardedCounter) Delete(key string) {
	s.getShard(key).Delete(key)
}

// Expire sets the expiration time of the counter with the given key.
// Returns true if the expiration time was set, false if the counter does not exist.
func (s *ShardedCounter) Expire(key string, ttl int64) bool {
	return s.getShard(key).Update(key, func(c *vals.Counter) bool {
		if c.IsExpired() {
			return false
		}
		c.Expire(ttl)
		return true
	})
}

// TTL returns the remaining time of the counter with the given key.
// Returns -1 and true if the counter has no expiration time.
// Returns -2 and false if the counter does not exist.
func (s *ShardedCounter) TTL(key string) (int64, bool) {
	shard := s.getShard(key)
	c, ok := shard.Get(key)
	if !ok || c.IsExpired() {
		return -2, false
	}
	if c.TTL() == 0 {
		return -1, true
	}
	return c.TTL(), true
}

// IncrBy increments the counter with the given key by alpha.
// Returns 0 and false if the counter does not exist or the increment fails due to limit.
func (s *ShardedCounter) IncrBy(key string, alpha int64) (int64, bool) {
	shard := s.getShard(key)
	c, ok := shard.Get(key)
	if !ok {
		return 0, false
	}
	if c.IsExpired() {
		shard.Delete(key)
		return 0, false
	}
	ok = c.IncrBy(alpha)
	if !ok {
		return 0, false
	}
	return c.GetValue(), true
}

// DecrBy decrements the counter with the given key by alpha.
// Returns 0 and false if the counter does not exist or the decrement fails.
func (s *ShardedCounter) DecrBy(key string, alpha int64) (int64, bool) {
	shard := s.getShard(key)
	c, exist := shard.Get(key)
	if !exist {
		return 0, false
	}
	if c.IsExpired() {
		shard.Delete(key)
		return 0, false
	}
	ok := c.DecrBy(alpha)
	if !ok {
		return 0, false
	}
	return c.GetValue(), true
}

// Reset resets the counter with the given key to 0.
// Returns false if the counter does not exist.
func (s *ShardedCounter) Reset(key string) bool {
	shard := s.getShard(key)
	c, ok := shard.Get(key)
	if !ok {
		return false
	}
	if c.IsExpired() {
		shard.Delete(key)
		return false
	}
	c.Reset()
	return true
}

// CleanExpired cleans expired counters.
func (s *ShardedCounter) CleanExpired() {
	for _, shard := range s.shards {
		shard.Clean(func(key string, counter *vals.Counter) bool {
			return counter.IsExpired()
		})
	}
}

// Flush removes all counters from the shards.
func (s *ShardedCounter) Flush() {
	for _, shard := range s.shards {
		shard.Flush()
	}
}
