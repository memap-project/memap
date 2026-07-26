package tsmap

import "sync"

type TypedSyncMap[K comparable, V any] struct {
	m sync.Map
}

func (tm *TypedSyncMap[K, V]) Store(key K, value V) {
	tm.m.Store(key, value)
}

func (tm *TypedSyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	actual, loaded := tm.m.LoadOrStore(key, value)
	return actual.(V), loaded
}

func (tm *TypedSyncMap[K, V]) Delete(key K) {
	tm.m.Delete(key)
}

func (tm *TypedSyncMap[K, V]) Load(key K) (V, bool) {
	actual, loaded := tm.m.Load(key)
	return actual.(V), loaded
}
