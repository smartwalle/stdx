package slicex

func Set[T any, K comparable, R any](slice []T, fn func(elem T) (K, R)) map[K]R {
	var n = len(slice)
	if n == 0 {
		return nil
	}
	var m = make(map[K]R, n)
	for _, elem := range slice {
		var k, v = fn(elem)
		m[k] = v
	}
	return m
}

func SetMatched[T any, K comparable, R any](slice []T, predicate func(elem T) bool, fn func(elem T) (K, R)) map[K]R {
	var n = len(slice)
	if n == 0 {
		return nil
	}
	var m = make(map[K]R, n)
	for _, elem := range slice {
		if !predicate(elem) {
			continue
		}
		var k, v = fn(elem)
		m[k] = v
	}
	return m
}
