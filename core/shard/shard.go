package shard

import (
	"sync"

	"github.com/dmi3midd/memap/core/item"
)

// Shard is a unit of data storage for sharded map.
type Shard struct {
	mu    sync.RWMutex
	items map[string]item.Item
}

// NewShard creates a new shard.
func NewShard() *Shard {
	return &Shard{
		mu:    sync.RWMutex{},
		items: make(map[string]item.Item),
	}
}

// Get retrieves an item from the shard.
func (s *Shard) Get(key string) (*item.Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.items[key]
	if !ok {
		return nil, false
	}
	return &value, true
}

// Set sets or updates an item in the shard.
func (s *Shard) Set(key string, item item.Item) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = item
}

// Delete deletes an item from the shard.
func (s *Shard) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, key)
}
