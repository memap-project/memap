package ns

import (
	"sync"

	"github.com/memap-project/memap/core/tsmap"
)

// NamespaceManager manages multiple isolated namespaces and a default namespace.
type NamespaceManager struct {
	mu         sync.RWMutex
	namespaces tsmap.TypedSyncMap[string, *Namespace]
	defaultNs  *Namespace
}

// NewNamespaceManager creates a new NamespaceManager with an initialized default namespace.
func NewNamespaceManager() *NamespaceManager {
	defaultNs := NewNamespace()
	return &NamespaceManager{
		mu:         sync.RWMutex{},
		namespaces: tsmap.TypedSyncMap[string, *Namespace]{},
		defaultNs:  defaultNs,
	}
}
