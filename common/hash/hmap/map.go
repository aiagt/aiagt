package hmap

import "github.com/aiagt/aiagt/common/hash"

type Map[E comparable, T any] map[E]T

func NewMap[E, K comparable, T hash.Comparable[K, E]](key K, vals ...T) Map[E, T] {
	m := make(Map[E, T], len(vals))
	for _, val := range vals {
		m[val.HashKey(key)] = val
	}

	return m
}

func NewMapWithValue[E, K comparable, V any, T hash.Comparable[K, E]](key K, vf func(T) V, vals ...T) Map[E, V] {
	m := make(Map[E, V], len(vals))
	for _, val := range vals {
		m[val.HashKey(key)] = vf(val)
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
