package utils

func Zero[T any]() T {
	var zero T

	return zero
}

func IsZero[T comparable](v T) bool {
	return v == Zero[T]()
}
