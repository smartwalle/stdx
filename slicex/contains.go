package slicex

// Contains 判断 slice 中是否包含指定元素
func Contains[T any](slice []T, fn func(elem T) bool) bool {
	for _, elem := range slice {
		if fn(elem) {
			return true
		}
	}
	return false
}
