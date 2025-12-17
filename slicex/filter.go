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

// OptimizedFilter 筛选出 slice 中满足指定条件的元素，返回的新 slice 与原始 slice 共享底层数组，
// 特别注意：本函数仅适用于过滤之后不再需要原始 slice 的场景。
func OptimizedFilter[T any](slice []T, fn func(elem T) bool) []T {
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

func FilterOne[T any](slice []T, fn func(a, b T) T) (r T) {
	if len(slice) == 0 {
		return r
	}
	r = slice[0]
	for _, elem := range slice[1:] {
		r = fn(r, elem)
	}
	return r
}

// Min 返回切片中最小的元素，如果切片为空则返回零值,
// fn 函数用于比较两个元素，返回 true 表示第一个元素小于第二个元素
func Min[T any](slice []T, fn func(a, b T) bool) (r T) {
	return FilterOne(slice, func(a, b T) T {
		if fn(a, b) {
			return a
		}
		return b
	})
}

// Max 返回切片中最大的元素，如果切片为空则返回零值,
// fn 函数用于比较两个元素，返回 true 表示第一个元素大于第二个元素
func Max[T any](slice []T, fn func(a, b T) bool) (r T) {
	return FilterOne(slice, func(a, b T) T {
		if fn(a, b) {
			return a
		}
		return b
	})
}
