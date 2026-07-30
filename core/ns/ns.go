package ns

import (
	"context"
	"sync"

	"github.com/dmi3midd/memap/core/shard/shhash"
	"github.com/dmi3midd/memap/core/shard/shmap"
)

// Namespace represents a namespace in the memap.
type Namespace struct {
	mu     sync.RWMutex
	shmap  *shmap.ShardedMap
	shhash *shhash.ShardedHash
	cancel context.CancelFunc
}

// NewNamespace creates a new namespace.
func NewNamespace() *Namespace {
	return &Namespace{
		mu:     sync.RWMutex{},
		shmap:  shmap.NewShardedMap(),
		shhash: shhash.NewShardedHash(),
		cancel: nil,
	}
}
