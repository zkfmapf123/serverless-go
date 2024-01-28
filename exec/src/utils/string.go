package utils

func InjectDefaultValueNotExist[T any](value T, isExist func(v T) bool, defaultValue T) T {
	if isExist(value) {
		return value
	}

	return defaultValue
}
