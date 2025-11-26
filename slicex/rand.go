package slicex

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func Rand[T any](slice []T) T {
	var n = len(slice)
	var elem T
	if n < 1 {
		return elem
	}
	return slice[random.Intn(n)]
}
