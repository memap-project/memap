package ns

import (
	"sync"

	"github.com/memap-project/memap/core/tsmap"
)

// NamespaceManager manages multiple namespaces.
// It contains a typed sync map of namespaces and a default namespace.
type NamespaceManager struct {
	mu         sync.RWMutex
	namespaces tsmap.TypedSyncMap[string, *Namespace]
	defaultNs  *Namespace
}

// NewNamespaceManager creates a new namespace manager.
func NewNamespaceManager() *NamespaceManager {
	defaultNs := NewNamespace()
	return &NamespaceManager{
		mu:         sync.RWMutex{},
		namespaces: tsmap.TypedSyncMap[string, *Namespace]{},
		defaultNs:  defaultNs,
	}
}
