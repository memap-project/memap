package ns

// Drop drops all keys and namespaces.
// func (nm *NamespaceManager) Drop() error {
// 	return nil
// }

// Flush flushes all keys in all namespaces.
// func (nm *NamespaceManager) Flush() error {
// 	return nil
// }

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
