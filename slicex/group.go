package slicex

// Group 对 slice 中的元素进行分组操作
func Group[T any, K comparable, R any](slice []T, fn func(elem T) (K, R)) map[K][]R {
	var newMap = make(map[K][]R)
	for _, elem := range slice {
		var key, value = fn(elem)
		newMap[key] = append(newMap[key], value)
	}
	return newMap
}

// GroupMatched 对 slice 中满足条件的元素进行分组操作
func GroupMatched[T any, K comparable, R any](slice []T, predicate func(elem T) bool, fn func(elem T) (K, R)) map[K][]R {
	var newMap = make(map[K][]R)
	for _, elem := range slice {
		if predicate(elem) {
			var key, value = fn(elem)
			newMap[key] = append(newMap[key], value)
		}
	}
	return newMap
}
