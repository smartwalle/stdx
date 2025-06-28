package slicex

// Each 对 slice 进行遍历
func Each[T any](slice []T, fn func(elem T)) {
	for _, elem := range slice {
		fn(elem)
	}
}

// ReverseEach 对 slice 进行反向遍历
func ReverseEach[T any](slice []T, fn func(elem T)) {
	for i := len(slice) - 1; i >= 0; i-- {
		fn(slice[i])
	}
}
