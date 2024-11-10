package ux

func Append[T any](slice []T, elems ...T) []T {
	el := len(elems)

	n := len(slice)
	c := cap(slice)

	if n+el > c {
		npq := make([]T, n, c*2+el)
		copy(npq, slice)
		slice = npq
	}
	slice = slice[0 : n+el]

	for j := 0; j < el; j++ {
		slice[n+j] = elems[j]
	}
	return slice
}
