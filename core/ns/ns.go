package ns

import (
	"sync"

	"github.com/memap-project/memap/core/shard/shcounter"
	"github.com/memap-project/memap/core/shard/shhash"
	"github.com/memap-project/memap/core/shard/shmap"
)

// Namespace represents a namespace in the memap.
type Namespace struct {
	mu        sync.RWMutex
	shmap     *shmap.ShardedMap
	shhash    *shhash.ShardedHash
	shcounter *shcounter.ShardedCounter
}

// NewNamespace creates a new namespace.
func NewNamespace() *Namespace {
	return &Namespace{
		mu:        sync.RWMutex{},
		shmap:     shmap.NewShardedMap(),
		shhash:    shhash.NewShardedHash(),
		shcounter: shcounter.NewShardedCounter(),
	}
}

// CleanExpired cleans expired keys.
func (n *Namespace) CleanExpired() {
	n.shmap.CleanExpired()
	n.shhash.CleanExpired()
}

// Flush removes all keys from the namespace.
func (n *Namespace) Flush() {
	n.shmap.Flush()
	n.shhash.Flush()
}
