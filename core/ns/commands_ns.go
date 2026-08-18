package ns

import "github.com/memap-project/memap/core/tsmap"

// GetNs retrieves a namespace by name.
func (nm *NamespaceManager) GetNs(name string) (*Namespace, bool) {
	ns, exists := nm.namespaces.Load(name)
	return ns, exists
}

// [Create] creates a new namespace by name.
// Returns [ErrNamespaceAlreadyExists] if the namespace already exists.
func (nm *NamespaceManager) Create(name string) error {
	ns := NewNamespace()
	_, loaded := nm.namespaces.LoadOrStore(name, ns)
	if loaded {
		return ErrNamespaceAlreadyExists
	}
	return nil
}

// [Drop] deletes a namespace by name.
func (nm *NamespaceManager) Drop(name string) error {
	nm.namespaces.Delete(name)
	return nil
}

// Erase drops all keys and namespaces.
func (nm *NamespaceManager) Erase() {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	nm.flushDefaultNs()
	nm.namespaces = tsmap.TypedSyncMap[string, *Namespace]{}
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

// CleanExpired cleans all expired keys in all namespaces.
func (nm *NamespaceManager) CleanExpired() {
	nm.cleanDefaultNs()
	nm.cleanCustomNamespaces()
}

// cleanDefaultNs cleans all expired keys in the default namespace.
func (nm *NamespaceManager) cleanDefaultNs() {
	nm.defaultNs.CleanExpired()
}

// cleanCustomNamespaces cleans all expired keys in all namespaces except default.
func (nm *NamespaceManager) cleanCustomNamespaces() {
	nm.namespaces.Range(func(k string, v *Namespace) bool {
		v.CleanExpired()
		return true
	})
}
