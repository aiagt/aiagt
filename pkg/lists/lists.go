package lists

func FlatMap[T, E any](list []T, m func(t T) []E) []E {
	var result []E

	for _, v := range list {
		result = append(result, m(v)...)
	}

	return result
}

func Map[T, E any](list []T, m func(T) E) []E {
	var result []E

	for _, v := range list {
		result = append(result, m(v))
	}

	return result
}

func Filter[T any](t []T, f func(T) bool) []T {
	var result []T

	for _, v := range t {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}
