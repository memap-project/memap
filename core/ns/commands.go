package ns

import "time"

// [CreateNs] creates a new namespace by name.
// Returns [ErrInvalidNamespace] if the namespace is "sys".
// Returns [ErrNamespaceAlreadyExists] if the namespace already exists.
func (nm *NamespaceManager) CreateNs(name string) error {
	nm.mu.Lock()
	defer nm.mu.Unlock()

	if name == "sys" {
		return ErrInvalidNamespace
	}

	if _, exists := nm.namespaces[name]; exists {
		return ErrNamespaceAlreadyExists
	}

	ns := NewNamespace()
	nm.namespaces[name] = ns
	return nil
}

// [DeleteNs] deletes a namespace by name.
// Returns [ErrInvalidNamespace] if the namespace is "sys".
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) DeleteNs(name string) error {
	nm.mu.Lock()
	defer nm.mu.Unlock()

	if name == "sys" {
		return ErrInvalidNamespace
	}

	if _, exists := nm.namespaces[name]; !exists {
		return ErrNamespaceNotFound
	}

	delete(nm.namespaces, name)
	return nil
}

// [Get] retrieves an item from the namespace by key.
// Returns [ErrKeyNotFound] if the key is not found.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) Get(ns, key string) (*Item, error) {
	nm.mu.RLock()
	defer nm.mu.RUnlock()

	if ns == "" {
		ns = "sys"
	}

	if n, exists := nm.namespaces[ns]; exists {
		item, exists := n.dict[key]
		if exists {
			if item.IsExpired() && !item.ExpiresAt.IsZero() {
				delete(n.dict, key)
				return nil, nil
			}
			return &item, nil
		}
		return nil, ErrKeyNotFound
	}
	return nil, ErrNamespaceNotFound
}

// [Set] stores an item in the namespace by key.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) Set(ns, key, value string, ttl time.Duration) error {
	nm.mu.Lock()
	defer nm.mu.Unlock()

	if ns == "" {
		ns = "sys"
	}

	item := Item{
		Value: value,
	}
	var zeroTtl time.Time
	if ttl == 0 {
		item.ExpiresAt = zeroTtl
	}

	if n, exists := nm.namespaces[ns]; exists {
		n.dict[key] = item
		return nil
	}
	return ErrNamespaceNotFound
}

// [Delete] removes an item from the namespace by key.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) Delete(ns, key string) error {
	nm.mu.Lock()
	defer nm.mu.Unlock()

	if ns == "" {
		ns = "sys"
	}

	if n, exists := nm.namespaces[ns]; exists {
		delete(n.dict, key)
		return nil
	}
	return ErrNamespaceNotFound
}
