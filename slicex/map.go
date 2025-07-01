package slicex

// Map 对 slice 中的元素进行转换
func Map[T any, N any](slice []T, fn func(elem T) N) []N {
	var n = len(slice)
	if n == 0 {
		return nil
	}
	var ns = make([]N, 0, n)
	for _, elem := range slice {
		ns = Append(ns, fn(elem))
	}
	return ns
}

// MapMatched 对 slice 中满足指定条件的元素进行转换
func MapMatched[T any, N any](slice []T, predicate func(elem T) bool, fn func(elem T) N) []N {
	var n = len(slice)
	if n == 0 {
		return nil
	}
	var ns = make([]N, 0, n)
	for _, elem := range slice {
		if predicate(elem) {
			ns = Append(ns, fn(elem))
		}
	}
	return ns
}
