package shcounter

import (
	"sync"
	"time"
)

type Counter struct {
	mu        sync.RWMutex
	value     int64
	limit     int64
	expiresAt int64
}

func NewCounter() *Counter {
	return &Counter{
		mu:        sync.RWMutex{},
		value:     0,
		limit:     0,
		expiresAt: 0,
	}
}

// IsExpired returns true if the hash is expired.
func (c *Counter) IsExpired() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.expiresAt == 0 {
		return false
	}
	return time.Now().Unix() > c.expiresAt
}

// Expire sets the expiration time of the counter.
// Returns false if the counter is already expired.
func (c *Counter) Expire(ttl int64) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.expiresAt > 0 || (c.expiresAt > 0 && time.Now().Unix() > c.expiresAt) {
		return false
	}
	c.expiresAt = time.Now().Unix() + ttl
	return true
}

// LeftTime returns the remaining time of the counter.
// Returns 0 if the counter has no expiration time.
func (c *Counter) LeftTime() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.expiresAt == 0 {
		return 0
	}
	return c.expiresAt - time.Now().Unix()
}

// SetLimit sets or updates the limit of the counter.
func (c *Counter) SetLimit(limit int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.limit = limit
}

// Val returns the current value of the counter.
func (c *Counter) Val() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

// Incr increments the counter by alpha, up to the limit if set.
// Returns true if the counter was incremented, false if the limit was reached.
func (c *Counter) Incr(alpha int64) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.limit > 0 && c.value+alpha > c.limit {
		return false
	}
	c.value += alpha
	return true
}

// Decr decrements the counter by alpha, down to 0 if the counter is already 0.
// Returns true if the counter was decremented, false otherwise.
func (c *Counter) Decr(alpha int64) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.value-alpha >= 0 {
		c.value -= alpha
		return true
	}
	return false
}
