package hash

type Comparable[K, T comparable] interface {
	HashKey(K) T
}
