package slicex

// Equals 比较两个 slice 是否相等
func Equals[T any](s1, s2 []T, fn func(a, b T) bool) bool {
	if len(s1) != len(s2) {
		return false
	}

	for idx := range s1 {
		if !fn(s1[idx], s2[idx]) {
			return false
		}
	}
	return true
}
