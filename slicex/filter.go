package slicex

// Filter 筛选出 slice 中满足条件的元素
func Filter[T any](slice []T, fn func(elem T) bool) []T {
	if len(slice) == 0 {
		return nil
	}
	var ns = make([]T, 0, len(slice))
	for _, elem := range slice {
		if fn(elem) {
			ns = Append(ns, elem)
		}
	}
	return ns
}
