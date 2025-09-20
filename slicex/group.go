package slicex

// Group 对 slice 中的元素进行分组操作
func Group[T any, K comparable](slice []T, fn func(elem T) K) map[K][]T {
	var newMap = make(map[K][]T)
	for _, elem := range slice {
		var key = fn(elem)
		newMap[key] = append(newMap[key], elem)
	}
	return newMap
}

// GroupMatched 对 slice 中满足条件的元素进行分组操作
func GroupMatched[T any, K comparable](slice []T, predicate func(elem T) bool, fn func(elem T) K) map[K][]T {
	var newMap = make(map[K][]T)
	for _, elem := range slice {
		if predicate(elem) {
			var key = fn(elem)
			newMap[key] = append(newMap[key], elem)
		}
	}
	return newMap
}
