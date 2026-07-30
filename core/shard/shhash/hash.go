package shhash

import (
	"maps"
	"sync"
)

type Hash struct {
	mu   sync.RWMutex
	hash map[string]string
}

func NewHash() *Hash {
	return &Hash{
		hash: make(map[string]string),
	}
}

func (h *Hash) GetMap() map[string]string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	cp := make(map[string]string, len(h.hash))
	maps.Copy(cp, h.hash)
	return cp
}

func (h *Hash) Get(key string) (string, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	value, ok := h.hash[key]
	if !ok {
		return "", false
	}
	return value, true
}

func (h *Hash) Set(key string, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.hash[key] = value
}

func (h *Hash) Delete(key string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.hash, key)
}
