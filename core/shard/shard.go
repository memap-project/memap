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
// Returns true if the item was retrieved, false if it was initialized.
func (s *Shard[V]) GetOrInit(key string, initFn func() V) (V, bool) {
	s.mu.RLock()
	val, ok := s.items[key]
	s.mu.RUnlock()
	if ok {
		return val, true
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if val, ok = s.items[key]; ok {
		return val, true
	}
	val = initFn()
	s.items[key] = val
	return val, false
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

// Update modifies an item in the shard under lock if it exists.
// Returns true if the item was found and updated.
func (s *Shard[V]) Update(key string, fn func(val V) bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.items[key]
	if !ok {
		return false
	}
	if fn(val) {
		s.items[key] = val
		return true
	}
	return false
}

// Clean cleans items from the shard based on a predicate function.
// The predicate function receives the key and value and should return true if the item should be deleted.
// This operation acquires a write lock for the entire duration.
func (s *Shard[V]) Clean(predicate func(key string, value V) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, v := range s.items {
		if predicate(k, v) {
			delete(s.items, k)
		}
	}
}

// Flush removes all items from the shard.
func (s *Shard[V]) Flush() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = make(map[string]V, 8)
}
