package slicex

// Compare 对 slice 的每个元素进行比较，返回其中的最大值或者最小值
func Compare[T any](slice []T, fn func(elem1, elem2 T) bool) T {
	var elem T
	if len(slice) == 0 {
		return elem
	}

	elem = slice[0]
	for i := 1; i < len(slice); i++ {
		if !fn(elem, slice[i]) {
			elem = slice[i]
		}
	}
	return elem
}

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Min 获取 slice 中的最小值
func Min[T Numeric](slice []T) T {
	return Compare(slice, func(elem1, elem2 T) bool {
		return elem1 < elem2
	})
}

// Max 获取 slice 中的最大值
func Max[T Numeric](slice []T) T {
	return Compare(slice, func(elem1, elem2 T) bool {
		return elem1 > elem2
	})
}
