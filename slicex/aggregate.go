package slicex

// Aggregate 对 slice 中的元素进行归约操作
func Aggregate[T any, R any](slice []T, initial R, fn func(seed R, elem T) R) R {
	var seed = initial
	for _, elem := range slice {
		seed = fn(seed, elem)
	}
	return seed
}
