package ns

import (
	"time"

	"github.com/dmi3midd/memap/core/item"
)

// [Get] retrieves an item from the namespace by key.
// Returns [ErrKeyNotFound] if the key is not found or expired.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) Get(ns, key string) (string, error) {
	if ns == "" {
		i, ok := nm.defaultNs.shmap.Get(key)
		if !ok {
			return "", ErrKeyNotFound
		}
		return i.Value, nil
	}

	n, exists := nm.GetNs(ns)
	if !exists {
		return "", ErrNamespaceNotFound
	}

	i, ok := n.shmap.Get(key)
	if !ok {
		return "", ErrKeyNotFound
	}
	return i.Value, nil
}

// [Set] stores an item in the namespace by key.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) Set(ns, key, value string, ttl int64) error {
	var t int64
	if ttl > 0 {
		t = time.Now().Unix() + ttl
	}

	item := item.Item{
		Value:     value,
		ExpiresAt: t,
	}

	if ns == "" {
		nm.defaultNs.shmap.Set(key, item)
		return nil
	}

	n, exists := nm.GetNs(ns)
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
		nm.defaultNs.shmap.Delete(key)
		return nil
	}

	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shmap.Delete(key)
	return nil
}

func (nm *NamespaceManager) Expire(ns, key string, ttl int64) error {
	if ns == "" {
		if ok := nm.defaultNs.shmap.Expire(key, ttl); !ok {
			return ErrKeyNotFound
		}
		return nil
	}

	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}

	if ok := n.shmap.Expire(key, ttl); !ok {
		return ErrKeyNotFound
	}
	return nil
}
