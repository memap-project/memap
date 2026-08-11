package ns

import (
	"sync"

	"github.com/dmi3midd/memap/core/shard/shhash"
	"github.com/dmi3midd/memap/core/shard/shmap"
)

// Namespace represents a namespace in the memap.
type Namespace struct {
	mu     sync.RWMutex
	shmap  *shmap.ShardedMap
	shhash *shhash.ShardedHash
}

// NewNamespace creates a new namespace.
func NewNamespace() *Namespace {
	return &Namespace{
		mu:     sync.RWMutex{},
		shmap:  shmap.NewShardedMap(),
		shhash: shhash.NewShardedHash(),
	}
}

// Clean cleans expired keys.
func (n *Namespace) Clean() {
	n.shmap.Clean()
	n.shhash.Clean()
}

// Flush removes all keys from the namespace.
func (n *Namespace) Flush() {
	n.shmap.Flush()
	n.shhash.Flush()
}
