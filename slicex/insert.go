package slicex

// Insert 插入元素到 slice
func Insert[T any](slice []T, idx int, elems ...T) []T {
	if len(elems) == 0 {
		return slice
	}
	if idx < 0 {
		idx = 0
	}
	if idx >= len(slice) {
		return Append(slice, elems...)
	}
	return Append(slice[:idx], Append(elems, slice[idx:]...)...)
}
