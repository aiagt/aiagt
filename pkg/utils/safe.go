package utils

func ValOf[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}

	return *v
}

func PtrOf[T any](v T) *T {
	return &v
}

func OPtrOf[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}

	return &v
}

func FirstResult[T any](t T, _ any) T {
	return t
}

func SecondResult[T any](_ any, t T, _ any) T {
	return t
}

func Or[T comparable](v1, v2 T) T {
	if IsZero(v1) {
		return v2
	}

	return v1
}
