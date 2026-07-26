package ns

import (
	"context"
	"sync"

	"github.com/dmi3midd/memap/core/tsmap"
)

type NamespaceManager struct {
	mu           sync.RWMutex
	ctx          context.Context
	cancel       context.CancelFunc
	namespaces   tsmap.TypedSyncMap[string, *Namespace]
	sysNamespace *Namespace
}

func NewNamespaceManager() *NamespaceManager {
	ctx, cancel := context.WithCancel(context.Background())
	systemNs := NewNamespace()
	return &NamespaceManager{
		mu:           sync.RWMutex{},
		ctx:          ctx,
		cancel:       cancel,
		namespaces:   tsmap.TypedSyncMap[string, *Namespace]{},
		sysNamespace: systemNs,
	}
}
