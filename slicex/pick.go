package slicex

// Pick 从 slice 中获取元素
func Pick[T any](slice []T, offset, limit int) []T {
	if offset < 0 || limit < 1 {
		return nil
	}
	var n = len(slice)
	if n == 0 {
		return nil
	}
	if offset >= n {
		return nil
	}
	var end = offset + limit
	if end > n {
		end = n
	}
	return slice[offset:end:end]
}
