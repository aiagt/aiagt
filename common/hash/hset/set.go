package hset

import "github.com/aiagt/aiagt/common/hash"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vals ...T) Set[T] {
	s := make(Set[T], len(vals))
	for _, v := range vals {
		s[v] = struct{}{}
	}

	return s
}

func NewSetWithKey[T, K comparable, E hash.Comparable[K, T]](key K, vals ...E) Set[T] {
	s := make(Set[T], len(vals))
	for _, v := range vals {
		s[v.HashKey(key)] = struct{}{}
	}

	return s
}

func NewSetWithFunc[T comparable, E any](fn func(E) T, vals ...E) Set[T] {
	s := make(Set[T], len(vals))
	for _, v := range vals {
		s[fn(v)] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(vals ...T) {
	for _, v := range vals {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Remove(v T) {
	if _, ok := s[v]; ok {
		delete(s, v)
	}
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) List() []T {
	result := make([]T, 0, len(s))
	for v := range s {
		result = append(result, v)
	}

	return result
}
