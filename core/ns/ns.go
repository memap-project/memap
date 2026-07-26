package ns

import (
	"context"
	"sync"

	"github.com/dmi3midd/memap/core/shard/shmap"
)

// Namespace represents a namespace in the memap system.
type Namespace struct {
	mu     sync.RWMutex
	shmap  *shmap.ShardedMap
	cancel context.CancelFunc
}

// NewNamespace creates a new namespace.
func NewNamespace() *Namespace {
	return &Namespace{
		mu:     sync.RWMutex{},
		shmap:  shmap.NewShardedMap(),
		cancel: nil,
	}
}
