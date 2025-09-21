package helpers

func Filter[T any](slc []T, fn func(value T) bool) []T {
	res := make([]T, len(slc))
	for _, value := range slc {
		if fn(value) {
			res = append(res, value)
		}
	}
	return res
}
