package ns

import (
	"context"
	"sync"

	"github.com/dmi3midd/memap/core/tsmap"
)

// NamespaceManager manages multiple namespaces.
// It contains a typed sync map of namespaces and a default namespace.
type NamespaceManager struct {
	mu         sync.RWMutex
	ctx        context.Context
	cancel     context.CancelFunc
	namespaces tsmap.TypedSyncMap[string, *Namespace]
	defaultNs  *Namespace
}

// NewNamespaceManager creates a new namespace manager.
func NewNamespaceManager() *NamespaceManager {
	ctx, cancel := context.WithCancel(context.Background())
	defaultNs := NewNamespace()
	return &NamespaceManager{
		mu:         sync.RWMutex{},
		ctx:        ctx,
		cancel:     cancel,
		namespaces: tsmap.TypedSyncMap[string, *Namespace]{},
		defaultNs:  defaultNs,
	}
}
