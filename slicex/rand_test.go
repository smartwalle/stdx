package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestRand(t *testing.T) {
	var values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		t.Log(slicex.Rand(values))
	}
}
