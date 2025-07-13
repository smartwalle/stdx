package slicex

// Unique 获取 slice 中的不重复元素
func Unique[T any, K comparable](slice []T, fn func(elem T) K) []T {
	var n = len(slice)
	if n < 2 {
		return slice
	}

	var keyMap = make(map[K]struct{}, n)
	var ns = make([]T, 0, n)

	for _, elem := range slice {
		var key = fn(elem)
		if _, ok := keyMap[key]; ok {
			continue
		}
		keyMap[key] = struct{}{}
		ns = Append(ns, elem)
	}
	return ns
}
