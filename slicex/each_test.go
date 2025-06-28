package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestEach(t *testing.T) {
	var s1 = []int{1, 2, 3, 4, 5}
	slicex.Each(s1, func(elem int) {
		t.Log(elem)
	})
}

func TestReverseEach(t *testing.T) {
	var s1 = []int{1, 2, 3, 4, 5}
	slicex.ReverseEach(s1, func(elem int) {
		t.Log(elem)
	})
}
