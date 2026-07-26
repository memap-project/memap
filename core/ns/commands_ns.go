package ns

// GetNs retrieves a namespace by name.
func (nm *NamespaceManager) GetNs(name string) (*Namespace, bool) {
	ns, exists := nm.namespaces.Load(name)
	return ns, exists
}

// [CreateNs] creates a new namespace by name.
// Returns [ErrNamespaceAlreadyExists] if the namespace already exists.
func (nm *NamespaceManager) CreateNs(name string) error {
	_, exists := nm.GetNs(name)
	if exists {
		return ErrNamespaceAlreadyExists
	}
	ns := NewNamespace()
	nm.namespaces.Store(name, ns)
	return nil
}

// [DeleteNs] deletes a namespace by name.
// Returns [ErrNamespaceNotFound] if the namespace does not exist.
func (nm *NamespaceManager) DeleteNs(name string) error {
	nm.namespaces.Delete(name)
	return nil
}
