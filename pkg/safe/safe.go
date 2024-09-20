package safe

func Value[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}

	return *v
}

func Pointer[T any](v T) *T {
	return &v
}

func OptionalPointer[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}

	return &v
}
