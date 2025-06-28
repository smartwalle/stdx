package stdx

// Append 添加元素到 Slice
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

// Split 将 Slice 转换成指定长度的二维 Slice
func Split[T any](slice []T, size int) [][]T {
	var nSlice [][]T
	if size < 1 {
		return nSlice
	}
	var remain = len(slice)
	for remain > 0 {
		var end = size
		if remain <= size {
			end = remain
		}
		nSlice = append(nSlice, slice[0:end])
		slice = slice[end:]
		remain = len(slice)
	}
	return nSlice
}

// Pick 从 Slice 中获取元素
func Pick[T any](slice []T, offset, limit int) []T {
	if offset < 0 || limit < 1 {
		return nil
	}
	var n = len(slice)
	if n < 1 {
		return nil
	}
	if offset >= n {
		return nil
	}
	var end = offset + limit
	if end > n {
		end = n
	}
	return slice[offset:end:end]
}
