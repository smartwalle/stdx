package slicex

// Append 添加元素到 slice
func Append[T any](slice []T, elems ...T) []T {
	n := len(slice)
	c := cap(slice)
	e := len(elems)

	if n+e > c {
		ns := make([]T, n, c*2+e)
		copy(ns, slice)
		slice = ns
	}
	slice = slice[0 : n+e]

	for i := 0; i < e; i++ {
		slice[n+i] = elems[i]
	}
	return slice
}
