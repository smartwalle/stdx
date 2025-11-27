package slicex

// Map 对 slice 中的元素进行转换
func Map[T any, R any](slice []T, fn func(elem T) R) []R {
	var n = len(slice)
	if n == 0 {
		return nil
	}
	var ns = make([]R, n)
	for idx, elem := range slice {
		ns[idx] = fn(elem)
	}
	return ns
}

// MapMatched 对 slice 中满足指定条件的元素进行转换
func MapMatched[T any, R any](slice []T, predicate func(elem T) bool, fn func(elem T) R) []R {
	var n = len(slice)
	if n == 0 {
		return nil
	}
	var ns = make([]R, 0, n)
	for _, elem := range slice {
		if predicate(elem) {
			ns = append(ns, fn(elem))
		}
	}
	return ns
}
