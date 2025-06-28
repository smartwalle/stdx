package slicex

// Split 将 slice 转换成指定长度的二维 slice
func Split[T any](slice []T, size int) [][]T {
	var ns [][]T
	if size < 1 {
		return ns
	}
	var n = len(slice)
	for n > 0 {
		var end = size
		if n <= size {
			end = n
		}
		ns = Append(ns, slice[0:end])
		slice = slice[end:]
		n = len(slice)
	}
	return ns
}
