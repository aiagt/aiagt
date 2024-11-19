package utils

func First[T any](list []T) T {
	var zero T
	if len(list) == 0 {
		return zero
	}

	return list[0]
}

func SafeSlice[E any, T []E](s T, start, end int) T {
	start = min(max(start, 0), len(s))
	end = min(max(start, 0), len(s))

	return s[start:end]
}

func SafeSubStr(s string, start, length int) string {
	if start < 0 || length < 0 {
		return ""
	}

	start = Min(start, len(s))
	end := Min(start+length, len(s))

	return s[start:end]
}
