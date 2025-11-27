package slicex

// Filter 筛选出 slice 中满足指定条件的元素
func Filter[T any](slice []T, fn func(elem T) bool) []T {
	if len(slice) == 0 {
		return nil
	}
	var ns = make([]T, 0, len(slice))
	for _, elem := range slice {
		if fn(elem) {
			ns = append(ns, elem)
		}
	}
	return ns
}

// FilterX 筛选出 slice 中满足指定条件的元素，返回的新 slice 与原始 slice 共享底层数组，
// 特别注意：本函数仅适用于过滤之后不再需要原始 slice 的场景。
func FilterX[T any](slice []T, fn func(elem T) bool) []T {
	if len(slice) == 0 {
		return nil
	}
	var ns = slice[:0]
	for _, elem := range slice {
		if fn(elem) {
			ns = append(ns, elem)
		}
	}
	return ns
}
