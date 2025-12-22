package slicex

// Reduce 对 slice 中的元素进行归约操作
func Reduce[T any, R any](slice []T, initial R, fn func(acc R, elem T) R) R {
	var accumulator = initial
	for _, elem := range slice {
		accumulator = fn(accumulator, elem)
	}
	return accumulator
}

// ReduceMatched 对 slice 中满足条件的元素进行归约操作
func ReduceMatched[T any, R any](slice []T, initial R, predicate func(elem T) bool, fn func(acc R, elem T) R) R {
	var accumulator = initial
	for _, elem := range slice {
		if predicate(elem) {
			accumulator = fn(accumulator, elem)
		}
	}
	return accumulator
}
