package ns

// Init initializes a counter in the specified namespace.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) Init(ns, key string, limit, ttl int64) error {
	if ns == "" {
		nm.defaultNs.shcounter.Init(key, limit, ttl)
		return nil
	}

	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shcounter.Init(key, limit, ttl)
	return nil
}

// SetLimit sets or updates the limit of a counter in the specified namespace.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
// Returns [ErrKeyNotFound] if the counter does not exist.
func (nm *NamespaceManager) SetLimit(ns, key string, limit int64) error {
	if ns == "" {
		nm.defaultNs.shcounter.SetLimit(key, limit)
		return nil
	}

	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	ok := n.shcounter.SetLimit(key, limit)
	if !ok {
		return ErrKeyNotFound
	}
	return nil
}

// CGet returns the value of a counter in the specified namespace.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
// Returns [ErrKeyNotFound] if the counter does not exist.
func (nm *NamespaceManager) CGet(ns, key string) (int64, error) {
	if ns == "" {
		count, ok := nm.defaultNs.shcounter.CGet(key)
		if !ok {
			return 0, ErrKeyNotFound
		}
		return count, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return 0, ErrNamespaceNotFound
	}
	count, ok := n.shcounter.CGet(key)
	if !ok {
		return 0, ErrKeyNotFound
	}
	return count, nil
}

// CDelete deletes a counter in the specified namespace.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) CDelete(ns, key string) error {
	if ns == "" {
		nm.defaultNs.shcounter.CDelete(key)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shcounter.CDelete(key)
	return nil
}

// CExpire sets the expiration time of a counter in the specified namespace.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
// Returns [ErrKeyNotFound] if the counter does not exist.
func (nm *NamespaceManager) CExpire(ns, key string, ttl int64) error {
	if ns == "" {
		ok := nm.defaultNs.shcounter.CExpire(key, ttl)
		if !ok {
			return ErrKeyNotFound
		}
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shcounter.CExpire(key, ttl)
	return nil
}

// CTTL returns the remaining time of a counter in the specified namespace.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
// Returns [ErrKeyNotFound] if the counter does not exist.
func (nm *NamespaceManager) CTTL(ns, key string) (int64, error) {
	if ns == "" {
		count, ok := nm.defaultNs.shcounter.CTTL(key)
		if !ok {
			return 0, ErrKeyNotFound
		}
		return count, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return 0, ErrNamespaceNotFound
	}
	count, ok := n.shcounter.CTTL(key)
	if !ok {
		return 0, ErrKeyNotFound
	}
	return count, nil
}

// Incr increments the counter in the specified namespace.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
// Returns [ErrKeyNotFound] if the counter does not exist.
func (nm *NamespaceManager) Incr(ns, key string, alpha int64) (int64, error) {
	if ns == "" {
		count, ok := nm.defaultNs.shcounter.Incr(key, alpha)
		if !ok {
			return 0, ErrKeyNotFound
		}
		return count, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return 0, ErrNamespaceNotFound
	}
	count, ok := n.shcounter.Incr(key, alpha)
	if !ok {
		return 0, ErrKeyNotFound
	}
	return count, nil
}

// Decr decrements the counter in the specified namespace.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
// Returns [ErrKeyNotFound] if the counter does not exist.
func (nm *NamespaceManager) Decr(ns, key string, alpha int64) (int64, error) {
	if ns == "" {
		count, ok := nm.defaultNs.shcounter.Decr(key, alpha)
		if !ok {
			return 0, ErrKeyNotFound
		}
		return count, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return 0, ErrNamespaceNotFound
	}
	count, ok := n.shcounter.Decr(key, alpha)
	if !ok {
		return 0, ErrKeyNotFound
	}
	return count, nil
}
