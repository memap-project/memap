package shard

import (
	"sync"
)

// Shard is a unit of data storage for sharded map.
type Shard[V any] struct {
	mu    sync.RWMutex
	items map[string]V
}

// NewShard creates a new shard.
func NewShard[V any]() *Shard[V] {
	return &Shard[V]{
		mu:    sync.RWMutex{},
		items: make(map[string]V, 8),
	}
}

// Get retrieves an item from the shard.
// Returns zero value and false if the key doesn't exist.
func (s *Shard[V]) Get(key string) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.items[key]
	if !ok {
		var zero V
		return zero, false
	}
	return value, true
}

// GetOrInit retrieves an item from the shard, or initializes and stores a new item if it does not exist.
func (s *Shard[V]) GetOrInit(key string, initFn func() V) V {
	s.mu.RLock()
	val, ok := s.items[key]
	s.mu.RUnlock()
	if ok {
		return val
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if val, ok = s.items[key]; ok {
		return val
	}
	val = initFn()
	s.items[key] = val
	return val
}

// Set sets or updates an item in the shard.
func (s *Shard[V]) Set(key string, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = value
}

// Delete deletes an item from the shard.
func (s *Shard[V]) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, key)
}
