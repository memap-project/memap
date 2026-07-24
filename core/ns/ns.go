package ns

import (
	"context"
	"sync"
	"time"
)

type Item struct {
	Value     string
	ExpiresAt time.Time
}

func (i *Item) IsExpired() bool {
	return time.Now().After(i.ExpiresAt)
}

type Namespace struct {
	mu     sync.RWMutex
	dict   map[string]Item
	cancel context.CancelFunc
}

func NewNamespace() *Namespace {
	return &Namespace{
		mu:   sync.RWMutex{},
		dict: make(map[string]Item),
	}
}
