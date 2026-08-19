package ns

import (
	"sync"

	"github.com/memap-project/memap/core/shard/shcounter"
	"github.com/memap-project/memap/core/shard/shhash"
	"github.com/memap-project/memap/core/shard/shmap"
)

// Namespace represents an isolated container storing maps, hashes, and counters.
type Namespace struct {
	mu        sync.RWMutex
	shmap     *shmap.ShardedMap
	shhash    *shhash.ShardedHash
	shcounter *shcounter.ShardedCounter
}

// NewNamespace creates a new Namespace with initialized storage components.
func NewNamespace() *Namespace {
	return &Namespace{
		mu:        sync.RWMutex{},
		shmap:     shmap.NewShardedMap(),
		shhash:    shhash.NewShardedHash(),
		shcounter: shcounter.NewShardedCounter(),
	}
}

// CleanExpired removes expired keys across all storage components in the namespace.
func (n *Namespace) CleanExpired() {
	n.shmap.CleanExpired()
	n.shhash.CleanExpired()
	n.shcounter.CleanExpired()
}

// Flush removes all keys across all storage components in the namespace.
func (n *Namespace) Flush() {
	n.shmap.Flush()
	n.shhash.Flush()
	n.shcounter.Flush()
}
