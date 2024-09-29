package hmap

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any](cap int) Map[K, V] {
	return make(Map[K, V], cap)
}

func FromMap[K comparable, V any](m map[K]V) Map[K, V] {
	return m
}

func FromMapEntries[T, K comparable, E, V any](m map[T]E, fn func(k T, v E) (K, V, bool)) Map[K, V] {
	result := NewMap[K, V](0)

	for k, v := range m {
		nk, nv, ok := fn(k, v)
		if ok {
			result[nk] = nv
		}
	}

	return result
}

func FromSliceEntries[K comparable, T, V any](vals []T, fn func(T) (K, V, bool)) Map[K, V] {
	m := make(Map[K, V], len(vals))

	for _, val := range vals {
		k, v, ok := fn(val)
		if ok {
			m[k] = v
		}
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
