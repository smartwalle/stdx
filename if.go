package stdx

func IF[T any](predicate bool, v1, v2 T) T {
	if predicate {
		return v1
	}
	return v2
}
