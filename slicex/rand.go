package slicex

import (
	"math/rand"
	"time"
)

var random *rand.Rand

func init() {
	var source = rand.NewSource(time.Now().UnixNano())
	random = rand.New(source)
}

func Rand[T any](slices []T) T {
	var n = len(slices)
	var elem T
	if n < 1 {
		return elem
	}
	return slices[random.Intn(n)]
}
