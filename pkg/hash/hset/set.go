package hset

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](cap int) Set[T] {
	return make(Set[T], cap)
}

func FromMap[K comparable, V any](m map[K]V) Set[K] {
	s := make(Set[K], len(m))
	for k, _ := range m {
		s[k] = struct{}{}
	}

	return s
}

func FromSlice[K comparable, T any](vals []T, fn func(T) K) Set[K] {
	s := make(Set[K])
	for _, val := range vals {
		s[fn(val)] = struct{}{}
	}

	return s
}

func (s Set[T]) Add(val T) {
	s[val] = struct{}{}
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := make(Set[T], max(len(s), len(other)))
	for v := range s {
		result.Add(v)
	}

	for v := range other {
		result.Add(v)
	}

	return result
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
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
