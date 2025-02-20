package utils

func Zero[T any]() T {
	var zero T

	return zero
}

func IsZero[T comparable](v T) bool {
	return v == Zero[T]()
}

func NonZero[T comparable](v T) bool {
	return v != Zero[T]()
}

func NonZeroAndNotEqual[T comparable](a, b T) bool {
	zero := Zero[T]()
	return a != zero && b != zero && a != b
}
