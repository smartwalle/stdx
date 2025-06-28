package slicex

// Map 对 slice 中的元素进行转换
func Map[T any, N any](slice []T, fn func(elem T) N) []N {
	n := len(slice)
	if n == 0 {
		return nil
	}
	var ns = make([]N, 0, n)
	for _, elem := range slice {
		ns = Append(ns, fn(elem))
	}
	return ns
}
