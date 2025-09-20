package slicex

// Aggregate 对 slice 中的元素进行聚合操作
func Aggregate[T any, R any](slice []T, initial R, fn func(acc R, elem T) R) R {
	var accumulator = initial
	for _, elem := range slice {
		accumulator = fn(accumulator, elem)
	}
	return accumulator
}

// AggregateMatched 对 slice 中满足条件的元素进行聚合操作
func AggregateMatched[T any, R any](slices []T, initial R, predicate func(elem T) bool, fn func(acc R, elem T) R) R {
	var accumulator = initial
	for _, elem := range slices {
		if predicate(elem) {
			accumulator = fn(accumulator, elem)
		}
	}
	return accumulator
}
