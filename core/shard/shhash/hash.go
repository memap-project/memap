package shhash

import (
	"maps"
	"sync"
	"time"
)

// Hash is a thread-safe map with RWMutex.
// It is used to store string key-value pairs.
type Hash struct {
	mu        sync.RWMutex
	hash      map[string]string
	expiresAt int64
}

// NewHash creates a new thread-safe map.
func NewHash() *Hash {
	return &Hash{
		hash:      make(map[string]string, 8),
		expiresAt: 0,
	}
}

// IsExpired returns true if the hash is expired.
func (h *Hash) IsExpired() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.expiresAt == 0 {
		return false
	}
	return time.Now().Unix() > h.expiresAt
}

// Expire sets the expiration time for the hash.
func (h *Hash) Expire(ttl int64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.expiresAt > 0 || (h.expiresAt > 0 && time.Now().Unix() > h.expiresAt) {
		return
	}
	h.expiresAt = time.Now().Unix() + ttl
}

func (h *Hash) LeftTime() int64 {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.expiresAt == 0 {
		return 0
	}
	return h.expiresAt - time.Now().Unix()
}

// GetCopy returns a copy of the hash map.
func (h *Hash) GetCopy() map[string]string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	cp := make(map[string]string, len(h.hash))
	maps.Copy(cp, h.hash)
	return cp
}

// Get retrieves a value for the given key.
// Returns empty string and false if the key does not exist.
func (h *Hash) Get(key string) (string, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	value, ok := h.hash[key]
	if !ok {
		return "", false
	}
	return value, true
}

// Set sets a value for the given key.
func (h *Hash) Set(key string, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.hash[key] = value
}

// Delete removes a key-value pair from the map.
func (h *Hash) Delete(key string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.hash, key)
}

// Len returns the number of key-value pairs in the map.
func (h *Hash) Len() int64 {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return int64(len(h.hash))
}

// Keys returns a slice of all keys in the hash.
func (h *Hash) Keys() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	keys := make([]string, 0, len(h.hash))
	for k := range h.hash {
		keys = append(keys, k)
	}
	return keys
}

func (h *Hash) Values() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	values := make([]string, 0, len(h.hash))
	for _, v := range h.hash {
		values = append(values, v)
	}
	return values
}
