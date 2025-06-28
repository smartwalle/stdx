package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestUnique(t *testing.T) {
	var tests = []struct {
		source []int
		expect []int
		fn     func(elem int) int
	}{
		{source: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}, expect: []int{1, 2, 3, 4, 5, 6}, fn: func(elem int) int {
			return elem
		}},
		{source: []int{3, 3, 3, 2, 2, 2, 1, 1, 1}, expect: []int{3, 2, 1}, fn: func(elem int) int {
			return elem
		}},
		{source: []int{1}, expect: []int{1}, fn: func(elem int) int {
			return elem
		}},
	}

	for _, test := range tests {
		var actual = slicex.Unique(test.source, test.fn)
		if !slicex.Equals(actual, test.expect, IntEqual) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}
