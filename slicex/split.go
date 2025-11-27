package slicex

// Split 将一维 slice 转换成指定长度的二维 slice
func Split[T any](slice []T, size int) [][]T {
	if size < 1 {
		return nil
	}
	var n = len(slice)
	var ns = make([][]T, 0, (n/size)+1)
	for n > 0 {
		var end = size
		if n <= size {
			end = n
		}
		ns = append(ns, slice[0:end])
		slice = slice[end:]
		n = len(slice)
	}
	return ns
}
