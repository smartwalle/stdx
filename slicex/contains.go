package slicex

// Contains 判断 slice 中是否包含满足指定条件的元素
func Contains[T any](slice []T, fn func(elem T) bool) bool {
	for _, elem := range slice {
		if fn(elem) {
			return true
		}
	}
	return false
}

// In 判断 slice 中是否包含指定元素
func In[T comparable](obj T, slice ...T) bool {
	if len(slice) < 1 {
		return false
	}
	for _, elem := range slice {
		if elem == obj {
			return true
		}
	}
	return false
}
