package helpers

func FindIndex[T any](list []T, fn func(val T) bool) int {
	index := -1
	for i, el := range list {
		if fn(el) {
			return i
		}
	}
	return index
}
