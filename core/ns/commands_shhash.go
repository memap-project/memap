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
func (nm *NamespaceManager) HSet(ns, key string, ttl int64) error {
	if ns == "" {
		nm.defaultNs.shhash.HSet(key, ttl)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shhash.HSet(key, ttl)
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

func (nm *NamespaceManager) HExpire(ns, key string, ttl int64) error {
	if ns == "" {
		nm.defaultNs.shhash.HExpire(key, ttl)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shhash.HExpire(key, ttl)
	return nil
}

// HTTL returns the time-to-live of the hash for the given key.
// Returns 0 if the key does not exist or has no TTL.
func (nm *NamespaceManager) HTTL(ns, key string) (int64, error) {
	if ns == "" {
		return nm.defaultNs.shhash.HTTL(key), nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return 0, ErrNamespaceNotFound
	}
	return n.shhash.HTTL(key), nil
}

func (nm *NamespaceManager) HExists(ns, key string) (bool, error) {
	if ns == "" {
		return nm.defaultNs.shhash.HExists(key), nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return false, ErrNamespaceNotFound
	}
	return n.shhash.HExists(key), nil
}

// HLEN returns the number of fields in the hash for the given key.
func (nm *NamespaceManager) HLen(ns, key string) (int64, error) {
	if ns == "" {
		return nm.defaultNs.shhash.HLen(key), nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return 0, ErrNamespaceNotFound
	}
	return n.shhash.HLen(key), nil
}

func (nm *NamespaceManager) HKeys(ns, key string) ([]string, error) {
	if ns == "" {
		return nm.defaultNs.shhash.HKeys(key), nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return []string{}, ErrNamespaceNotFound
	}
	return n.shhash.HKeys(key), nil
}

// HValues returns the values of the hash for the given key.
func (nm *NamespaceManager) HValues(ns, key string) ([]string, error) {
	if ns == "" {
		return nm.defaultNs.shhash.HValues(key), nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return nil, ErrNamespaceNotFound
	}
	return n.shhash.HValues(key), nil
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
