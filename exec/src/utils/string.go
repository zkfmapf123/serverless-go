package utils

import "fmt"

func InjectDefaultValueNotExist[T any](value T, isExist func(v T) bool, defaultValue T) T {
	if isExist(value) {
		return value
	}

	return defaultValue
}

func Concat(prev string, next ...string) string {

	answer := prev

	for _, n := range next {
		answer = fmt.Sprintf("%s%s", answer, n)
	}

	return answer
}
