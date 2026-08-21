package ns

func (nm *NamespaceManager) BInit(ns, key string, capacity, ttl int64) error {
	if ns == "" {
		ok := nm.defaultNs.shrbuffer.Init(key, capacity, ttl)
		if !ok {
			return ErrKeyAlreadyExists
		}
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	ok := n.shrbuffer.Init(key, capacity, ttl)
	if !ok {
		return ErrKeyAlreadyExists
	}
	return nil
}

func (nm *NamespaceManager) BPush(ns, key, value string) error {
	if ns == "" {
		ok := nm.defaultNs.shrbuffer.Push(key, value)
		if !ok {
			return ErrKeyNotFound
		}
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	ok := n.shrbuffer.Push(key, value)
	if !ok {
		return ErrKeyNotFound
	}
	return nil
}

func (nm *NamespaceManager) BPop(ns, key string) (string, error) {
	if ns == "" {
		value, ok := nm.defaultNs.shrbuffer.Pop(key)
		if !ok {
			return value, ErrKeyNotFound
		}
		return value, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return "", ErrNamespaceNotFound
	}
	value, ok := n.shrbuffer.Pop(key)
	if !ok {
		return value, ErrKeyNotFound
	}
	return value, nil
}

func (nm *NamespaceManager) BAt(ns, key string, index int64) (string, error) {
	if ns == "" {
		value, ok := nm.defaultNs.shrbuffer.At(key, index)
		if !ok {
			return value, ErrKeyNotFound
		}
		return value, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return "", ErrNamespaceNotFound
	}
	value, ok := n.shrbuffer.At(key, index)
	if !ok {
		return value, ErrKeyNotFound
	}
	return value, nil
}

func (nm *NamespaceManager) BSlice(ns, key string) ([]string, error) {
	if ns == "" {
		values, ok := nm.defaultNs.shrbuffer.Slice(key)
		if !ok {
			return nil, ErrKeyNotFound
		}
		return values, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return nil, ErrNamespaceNotFound
	}
	values, ok := n.shrbuffer.Slice(key)
	if !ok {
		return nil, ErrKeyNotFound
	}
	return values, nil
}

func (nm *NamespaceManager) BPeek(ns, key string) (string, error) {
	if ns == "" {
		value, ok := nm.defaultNs.shrbuffer.Peek(key)
		if !ok {
			return value, ErrKeyNotFound
		}
		return value, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return "", ErrNamespaceNotFound
	}
	value, ok := n.shrbuffer.Peek(key)
	if !ok {
		return value, ErrKeyNotFound
	}
	return value, nil
}

func (nm *NamespaceManager) BBack(ns, key string) (string, error) {
	if ns == "" {
		value, ok := nm.defaultNs.shrbuffer.Back(key)
		if !ok {
			return value, ErrKeyNotFound
		}
		return value, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return "", ErrNamespaceNotFound
	}
	value, ok := n.shrbuffer.Back(key)
	if !ok {
		return value, ErrKeyNotFound
	}
	return value, nil
}

func (nm *NamespaceManager) BCap(ns, key string) (int64, error) {
	if ns == "" {
		cap, ok := nm.defaultNs.shrbuffer.Cap(key)
		if !ok {
			return cap, ErrKeyNotFound
		}
		return cap, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return 0, ErrNamespaceNotFound
	}
	cap, ok := n.shrbuffer.Cap(key)
	if !ok {
		return cap, ErrKeyNotFound
	}
	return cap, nil
}

func (nm *NamespaceManager) BLen(ns, key string) (int64, error) {
	if ns == "" {
		len, ok := nm.defaultNs.shrbuffer.Len(key)
		if !ok {
			return len, ErrKeyNotFound
		}
		return len, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return 0, ErrNamespaceNotFound
	}
	len, ok := n.shrbuffer.Len(key)
	if !ok {
		return len, ErrKeyNotFound
	}
	return len, nil
}

func (nm *NamespaceManager) BReset(ns, key string) error {
	if ns == "" {
		nm.defaultNs.shrbuffer.Reset(key)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shrbuffer.Reset(key)
	return nil
}

func (nm *NamespaceManager) BDel(ns, key string) error {
	if ns == "" {
		nm.defaultNs.shrbuffer.Delete(key)
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	n.shrbuffer.Delete(key)
	return nil
}

func (nm *NamespaceManager) BExpire(ns, key string, ttl int64) error {
	if ns == "" {
		ok := nm.defaultNs.shrbuffer.Expire(key, ttl)
		if !ok {
			return ErrKeyNotFound
		}
		return nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return ErrNamespaceNotFound
	}
	ok := n.shrbuffer.Expire(key, ttl)
	if !ok {
		return ErrKeyNotFound
	}
	return nil
}

func (nm *NamespaceManager) BTTL(ns, key string) (int64, error) {
	if ns == "" {
		ttl, ok := nm.defaultNs.shrbuffer.TTL(key)
		if !ok {
			return ttl, ErrKeyNotFound
		}
		return ttl, nil
	}
	n, exists := nm.GetNs(ns)
	if !exists {
		return 0, ErrNamespaceNotFound
	}
	ttl, ok := n.shrbuffer.TTL(key)
	if !ok {
		return ttl, ErrKeyNotFound
	}
	return ttl, nil
}
