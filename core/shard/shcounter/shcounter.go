package shcounter

import (
	"hash/fnv"

	"github.com/memap-project/memap/core/shard"
)

type ShardedCounter struct {
	shardCount uint8
	shards     []*shard.Shard[*Counter]
}

func NewShardedCounter() *ShardedCounter {
	shards := make([]*shard.Shard[*Counter], 8)
	for i := range shards {
		shards[i] = shard.NewShard[*Counter]()
	}
	return &ShardedCounter{
		shardCount: 8,
		shards:     shards,
	}
}

// getShard returns the shard that the key belongs to.
func (s *ShardedCounter) getShard(key string) *shard.Shard[*Counter] {
	h := fnv.New32a()
	h.Write([]byte(key))
	idx := h.Sum32() & uint32(s.shardCount-1)
	return s.shards[idx]
}

// Init initializes a counter with the given key, limit, and ttl.
// Returns true if the counter was initialized, false if it already existed.
func (s *ShardedCounter) Init(key string, limit int64, ttl int64) bool {
	c, ok := s.getShard(key).GetOrInit(key, NewCounter)
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

// CGet returns the value and expiration status of the counter with the given key.
// Returns 0 and false if the counter does not exist.
func (s *ShardedCounter) CGet(key string) (int64, bool) {
	shard := s.getShard(key)
	c, ok := shard.Get(key)
	if !ok {
		return 0, false
	}
	if c.IsExpired() {
		shard.Delete(key)
		return 0, false
	}
	return c.Val(), true
}

// CDelete deletes the counter with the given key.
func (s *ShardedCounter) CDelete(key string) {
	s.getShard(key).Delete(key)
}

// CExpire sets the expiration time of the counter with the given key.
// Returns true if the expiration time was set, false if the counter does not exist or the expiration time already set.
func (s *ShardedCounter) CExpire(key string, ttl int64) bool {
	c, ok := s.getShard(key).Get(key)
	if !ok {
		return false
	}
	return c.Expire(ttl)
}

// CTTL returns the remaining time of the counter with the given key.
// Returns 0 if the counter does not have an expiration time.
// Returns 0 and false if the counter does not exist.
func (s *ShardedCounter) CTTL(key string) (int64, bool) {
	shard := s.getShard(key)
	c, ok := shard.Get(key)
	if !ok {
		return 0, false
	}
	if c.IsExpired() {
		shard.Delete(key)
		return 0, false
	}
	return c.LeftTime(), true
}

// Incr increments the counter with the given key by alpha.
// Returns 0 and false if the counter does not exist or the increment fails due to limit.
func (s *ShardedCounter) Incr(key string, alpha int64) (int64, bool) {
	shard := s.getShard(key)
	c, ok := shard.Get(key)
	if !ok {
		return 0, false
	}
	if c.IsExpired() {
		shard.Delete(key)
		return 0, false
	}
	ok = c.Incr(alpha)
	if !ok {
		return 0, false
	}
	return c.Val(), true
}

// Decr decrements the counter with the given key by alpha.
// Returns 0 and false if the counter does not exist or the decrement fails.
func (s *ShardedCounter) Decr(key string, alpha int64) (int64, bool) {
	shard := s.getShard(key)
	c, exist := shard.Get(key)
	if !exist {
		return 0, false
	}
	if c.IsExpired() {
		shard.Delete(key)
		return 0, false
	}
	ok := c.Decr(alpha)
	if !ok {
		return 0, false
	}
	return c.Val(), true
}

func (s *ShardedCounter) CleanExpired() {
	for _, shard := range s.shards {
		shard.Clean(func(key string, counter *Counter) bool {
			return counter.IsExpired()
		})
	}
}

func (s *ShardedCounter) Flush() {
	for _, shard := range s.shards {
		shard.Flush()
	}
}
