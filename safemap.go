package safemap

import "sync"

type SafeMap[K comparable, V any] interface {
	Delete(key K)
	GetOrZero(key K) V
	Get(key K) (V, bool)
	GetAndDelete(key K) (V, bool)
	GetOrSet(key K, value V) (V, bool)
	Range(f func(key K, value V) bool)
	Set(key K, value V)
	Size() int
}

type safeMap[K comparable, V any] struct {
	m   sync.Map
	len int
}

func New[K comparable, V any]() SafeMap[K, V] {
	return &safeMap[K, V]{}
}

func (m *safeMap[K, V]) Delete(key K) {
	m.m.Delete(key)
	m.len--
}

func (m *safeMap[K, V]) GetOrZero(key K) V {
	v, ok := m.m.Load(key)
	if !ok {
		var zero V
		return zero
	}
	return v.(V)
}

func (m *safeMap[K, V]) Get(key K) (V, bool) {
	v, ok := m.m.Load(key)
	if !ok {
		var zero V
		return zero, ok
	}
	return v.(V), ok
}

func (m *safeMap[K, V]) GetAndDelete(key K) (V, bool) {
	v, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		var zero V
		return zero, loaded
	}
	m.len--
	return v.(V), loaded
}

func (m *safeMap[K, V]) GetOrSet(key K, value V) (V, bool) {
	a, loaded := m.m.LoadOrStore(key, value)
	if !loaded {
		m.len++
	}
	return a.(V), loaded
}

func (m *safeMap[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool { return f(key.(K), value.(V)) })
}

func (m *safeMap[K, V]) Set(key K, value V) {
	m.m.Store(key, value)
	m.len++
}

func (m *safeMap[K, V]) Size() int {
	return m.len
}
