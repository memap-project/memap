package ns

import (
	"context"
	"sync"
)

type NamespaceManager struct {
	mu         sync.RWMutex
	namespaces map[string]*Namespace
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewNamespaceManager() *NamespaceManager {
	ctx, cancel := context.WithCancel(context.Background())
	systemNs := NewNamespace()
	return &NamespaceManager{
		mu:     sync.RWMutex{},
		ctx:    ctx,
		cancel: cancel,
		namespaces: map[string]*Namespace{
			"sys": systemNs,
		},
	}
}
