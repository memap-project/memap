package ns

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
