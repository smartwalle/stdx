package slicex

// Pick 从 slice 中获取元素
func Pick[T any](slice []T, offset, limit int) []T {
	var n = len(slice)
	if offset < 0 || limit < 1 || offset >= n {
		return nil
	}
	var end = offset + limit
	if end > n {
		end = n
	}
	return slice[offset:end:end]
}
