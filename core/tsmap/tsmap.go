package tsmap

import "sync"

// TypedSyncMap is a generic sync map that stores values of a specific type.
type TypedSyncMap[K comparable, V any] struct {
	m sync.Map
}

func (tm *TypedSyncMap[K, V]) Load(key K) (V, bool) {
	actual, loaded := tm.m.Load(key)
	if !loaded {
		var zero V
		return zero, false
	}
	return actual.(V), true
}

func (tm *TypedSyncMap[K, V]) Store(key K, value V) {
	tm.m.Store(key, value)
}

func (tm *TypedSyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	actual, loaded := tm.m.LoadOrStore(key, value)
	if !loaded {
		var zero V
		return zero, false
	}
	return actual.(V), true
}

func (tm *TypedSyncMap[K, V]) Delete(key K) {
	tm.m.Delete(key)
}
