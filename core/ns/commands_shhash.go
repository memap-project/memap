package ns

// HGet retrieves a hash from the namespace by key.
// Returns [ErrKeyNotFound] if the hash is not found.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HGet(ns, key string) (map[string]string, error) {
	if ns == "" {
		v, ok := nm.defaultNs.shhash.Get(key)
		if !ok {
			return v, ErrKeyNotFound
		}
		return v, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return map[string]string{}, ErrNamespaceNotFound
	}
	v, ok := n.shhash.Get(key)
	if !ok {
		return v, ErrKeyNotFound
	}
	return v, nil
}

// HSet creates an empty hash for the given key.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HSet(ns, key string, ttl int64) error {
	if ns == "" {
		nm.defaultNs.shhash.Set(key, ttl)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shhash.Set(key, ttl)
	return nil
}

// HDelete deletes a hash from the namespace by key.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HDel(ns, key string) error {
	if ns == "" {
		nm.defaultNs.shhash.Delete(key)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shhash.Delete(key)
	return nil
}

func (nm *NamespaceManager) HExpire(ns, key string, ttl int64) error {
	if ns == "" {
		ok := nm.defaultNs.shhash.Expire(key, ttl)
		if !ok {
			return ErrKeyNotFound
		}
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	ok := n.shhash.Expire(key, ttl)
	if !ok {
		return ErrKeyNotFound
	}
	return nil
}

// HTTL returns the time-to-live of the hash for the given key.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
// Returns -1 and false if the key has no expiration time.
// Returns -2 and false if the key does not exist.
func (nm *NamespaceManager) HTTL(ns, key string) (int64, error) {
	if ns == "" {
		return nm.defaultNs.shhash.TTL(key), nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return 0, ErrNamespaceNotFound
	}
	return n.shhash.TTL(key), nil
}

func (nm *NamespaceManager) HExists(ns, key string) (bool, error) {
	if ns == "" {
		return nm.defaultNs.shhash.Exists(key), nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return false, ErrNamespaceNotFound
	}
	return n.shhash.Exists(key), nil
}

// HLEN returns the number of fields in the hash for the given key.
func (nm *NamespaceManager) HLen(ns, key string) (int64, error) {
	if ns == "" {
		len, ok := nm.defaultNs.shhash.Len(key)
		if !ok {
			return len, ErrKeyNotFound
		}
		return len, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return 0, ErrNamespaceNotFound
	}
	len, ok := n.shhash.Len(key)
	if !ok {
		return len, ErrKeyNotFound
	}
	return len, nil
}

func (nm *NamespaceManager) HKeys(ns, key string) ([]string, error) {
	if ns == "" {
		keys, ok := nm.defaultNs.shhash.Keys(key)
		if !ok {
			return keys, ErrKeyNotFound
		}
		return keys, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return []string{}, ErrNamespaceNotFound
	}
	keys, ok := n.shhash.Keys(key)
	if !ok {
		return keys, ErrKeyNotFound
	}
	return keys, nil
}

// HValues returns the values of the hash for the given key.
func (nm *NamespaceManager) HValues(ns, key string) ([]string, error) {
	if ns == "" {
		values, ok := nm.defaultNs.shhash.Values(key)
		if !ok {
			return values, ErrKeyNotFound
		}
		return values, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return []string{}, ErrNamespaceNotFound
	}
	values, ok := n.shhash.Values(key)
	if !ok {
		return values, ErrKeyNotFound
	}
	return values, nil
}

// HFGet retrieves a field from the hash for the given key and field.
// Returns [ErrKeyNotFound] if the key or field does not exist.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HFGet(ns, key, field string) (string, error) {
	if ns == "" {
		v, ok := nm.defaultNs.shhash.GetField(key, field)
		if !ok {
			return v, ErrKeyNotFound
		}
		return v, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return "", ErrNamespaceNotFound
	}
	v, ok := n.shhash.GetField(key, field)
	if !ok {
		return v, ErrKeyNotFound
	}
	return v, nil
}

// HFSet sets a field in the hash for the given key and field.
// Creates hash if not exists.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HFSet(ns, key, field, value string) error {
	if ns == "" {
		nm.defaultNs.shhash.SetField(key, field, value)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shhash.SetField(key, field, value)
	return nil
}

// HFDelete deletes a field from the hash for the given key and field.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) HFDel(ns, key, field string) error {
	if ns == "" {
		nm.defaultNs.shhash.DeleteField(key, field)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shhash.DeleteField(key, field)
	return nil
}
