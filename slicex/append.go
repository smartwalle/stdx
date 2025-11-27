package slicex

// Append 追加元素到 slice，适用于没有明确指定 slice 容量的场景
func Append[T any](slice []T, elems ...T) []T {
	var n = len(slice)
	var c = cap(slice)
	var e = len(elems)

	if n+e > c {
		var ns = make([]T, n, c*2+e)
		copy(ns, slice)
		slice = ns
	}
	slice = slice[0 : n+e]

	for i := 0; i < e; i++ {
		slice[n+i] = elems[i]
	}
	return slice
}
