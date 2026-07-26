package ns

import (
	"context"
	"sync"

	"github.com/dmi3midd/memap/core/shard/shmap"
)

type Namespace struct {
	mu     sync.RWMutex
	shmap  *shmap.ShardedMap
	cancel context.CancelFunc
}

func NewNamespace() *Namespace {
	return &Namespace{
		mu:     sync.RWMutex{},
		shmap:  shmap.NewShardedMap(),
		cancel: nil,
	}
}
