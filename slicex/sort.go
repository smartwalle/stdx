package slicex

import (
	"sort"
)

// Sort 对 slice 中的元素进行排序,
// cmp 回调函数返回负数表示升序，返回正数表示降序，返回0表示相等
func Sort[T any](slice []T, cmps ...func(a, b T) int) {
	sort.Slice(slice, func(i, j int) bool {
		for _, cmp := range cmps {
			var v = cmp(slice[i], slice[j])
			if v < 0 {
				return true
			}
			if v > 0 {
				return false
			}
		}
		return false
	})
}
