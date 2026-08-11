package ns

// Erase drops all keys and namespaces.
func (nm *NamespaceManager) Erase() {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	nm.Flush()
	nm.namespaces.Range(func(k string, _ *Namespace) bool {
		nm.namespaces.Delete(k)
		return true
	})
}

// Flush flushes all keys in all namespaces.
func (nm *NamespaceManager) Flush() {
	nm.flushDefaultNs()
	nm.flushCustomNamespaces()
}

func (nm *NamespaceManager) flushDefaultNs() {
	nm.defaultNs.Flush()
}

func (nm *NamespaceManager) flushCustomNamespaces() {
	nm.namespaces.Range(func(k string, v *Namespace) bool {
		v.Flush()
		return true
	})
}

// Clean cleans all expired keys in all namespaces.
func (nm *NamespaceManager) Clean() {
	nm.cleanDefaultNs()
	nm.cleanCustomNamespaces()
}

// cleanDefaultNs cleans all expired keys in the default namespace.
func (nm *NamespaceManager) cleanDefaultNs() {
	nm.defaultNs.Clean()
}

// cleanCustomNamespaces cleans all expired keys in all namespaces except default.
func (nm *NamespaceManager) cleanCustomNamespaces() {
	nm.namespaces.Range(func(k string, v *Namespace) bool {
		v.Clean()
		return true
	})
}
