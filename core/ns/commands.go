package ns

import (
	"time"

	"github.com/dmi3midd/memap/core/item"
)

func (nm *NamespaceManager) getNs(name string) (*Namespace, bool) {
	ns, exists := nm.namespaces.Load(name)
	return ns, exists
}

// [CreateNs] creates a new namespace by name.
// Returns [ErrNamespaceAlreadyExists] if the namespace already exists.
func (nm *NamespaceManager) CreateNs(name string) error {
	_, exists := nm.getNs(name)
	if exists {
		return ErrNamespaceAlreadyExists
	}
	ns := NewNamespace()
	nm.namespaces.Store(name, ns)
	return nil
}

// [DeleteNs] deletes a namespace by name.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) DeleteNs(name string) error {
	nm.namespaces.Delete(name)
	return nil
}

// [Get] retrieves an item from the namespace by key.
// Returns [ErrKeyNotFound] if the key is not found or expired.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) Get(ns, key string) (*item.Item, error) {
	if ns == "" {
		item, ok := nm.sysNamespace.shmap.Get(key)
		if !ok {
			return nil, ErrKeyNotFound
		}
		return item, nil
	}

	n, exists := nm.getNs(ns)
	if !exists {
		return nil, ErrNamespaceNotFound
	}

	item, ok := n.shmap.Get(key)
	if !ok {
		return nil, ErrKeyNotFound
	}
	if item.IsExpired() {
		n.shmap.Delete(key)
		return nil, ErrKeyNotFound
	}
	return item, nil
}

// [Set] stores an item in the namespace by key.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) Set(ns, key, value string, ttl time.Duration) error {
	var t time.Time
	if ttl > 0 {
		t = time.Now().Add(ttl * time.Second)
	}

	item := item.Item{
		Value:     value,
		ExpiresAt: t,
	}

	if ns == "" {
		nm.sysNamespace.shmap.Set(key, item)
		return nil
	}

	n, exists := nm.getNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}

	n.shmap.Set(key, item)
	return nil
}

// [Delete] removes an item from the namespace by key.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) Delete(ns, key string) error {
	if ns == "" {
		nm.sysNamespace.shmap.Delete(key)
		return nil
	}

	n, exists := nm.getNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shmap.Delete(key)
	return nil
}
