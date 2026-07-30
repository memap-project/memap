package ns

// HGet retrieves a hash from the namespace by key.
// Returns [ErrKeyNotFound] if the hash is not found.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HGet(ns, key string) (map[string]string, error) {
	if ns == "" {
		v, ok := nm.defaultNs.shhash.HGet(key)
		if !ok {
			return nil, ErrKeyNotFound
		}
		return v, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return nil, ErrNamespaceNotFound
	}
	v, ok := n.shhash.HGet(key)
	if !ok {
		return nil, ErrKeyNotFound
	}
	return v, nil
}

// HSet creates an empty hash for the given key.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HSet(ns, key string) error {
	if ns == "" {
		nm.defaultNs.shhash.HSet(key)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shhash.HSet(key)
	return nil
}

// HDelete deletes a hash from the namespace by key.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HDelete(ns, key string) error {
	if ns == "" {
		nm.defaultNs.shhash.HDelete(key)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shhash.HDelete(key)
	return nil
}

// HFGet retrieves a field from the hash for the given key and field.
// Returns [ErrKeyNotFound] if the key or field does not exist.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HFGet(ns, key, field string) (string, error) {
	if ns == "" {
		v, ok := nm.defaultNs.shhash.HFGet(key, field)
		if !ok {
			return "", ErrKeyNotFound
		}
		return v, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return "", ErrNamespaceNotFound
	}
	v, ok := n.shhash.HFGet(key, field)
	if !ok {
		return "", ErrKeyNotFound
	}
	return v, nil
}

// HFSet sets a field in the hash for the given key and field.
// Creates hash if not exists.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HFSet(ns, key, field, value string) error {
	if ns == "" {
		nm.defaultNs.shhash.HFSet(key, field, value)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shhash.HFSet(key, field, value)
	return nil
}

// HFDelete deletes a field from the hash for the given key and field.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HFDelete(ns, key, field string) error {
	if ns == "" {
		nm.defaultNs.shhash.HFDelete(key, field)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shhash.HFDelete(key, field)
	return nil
}
