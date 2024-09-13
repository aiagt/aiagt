package hmap

import "github.com/aiagt/aiagt/common/hash"

type Map[E comparable, T any] map[E]T

func NewMap[E, K comparable, T hash.Comparable[K, E]](vals []T, key K) Map[E, T] {
	m := make(Map[E, T], len(vals))
	for _, val := range vals {
		m[val.HashKey(key)] = val
	}

	return m
}

func NewMapWithValueFunc[E, K comparable, V any, T hash.Comparable[K, E]](vals []T, key K, vf func(T) V) Map[E, V] {
	m := make(Map[E, V], len(vals))
	for _, val := range vals {
		m[val.HashKey(key)] = vf(val)
	}

	return m
}

func NewMapWithKeyFunc[E comparable, T any](vals []T, kf func(T) E) Map[E, T] {
	m := make(Map[E, T], len(vals))
	for _, val := range vals {
		m[kf(val)] = val
	}

	return m
}

func NewMapWithFuncs[E comparable, V, T any](vals []T, kf func(T) E, vf func(T) V) Map[E, V] {
	m := make(Map[E, V], len(vals))
	for _, val := range vals {
		m[kf(val)] = vf(val)
	}
	return m
}

func (m Map[E, T]) Keys() []E {
	var keys []E
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func (m Map[E, T]) Values() []T {
	var values []T
	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func (m Map[E, T]) Get(key E) T {
	val, ok := m[key]
	if !ok {
		var zero T
		return zero
	}
	return val
}
