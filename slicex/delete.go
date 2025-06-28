package slicex

// Delete 删除指定索引的元素
func Delete[T any](slice []T, idx int) []T {
	if idx < 0 || idx >= len(slice) {
		return slice
	}
	return Append(slice[:idx], slice[idx+1:]...)
}
