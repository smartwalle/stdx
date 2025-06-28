package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestFilter(t *testing.T) {
	var tests = []struct {
		source []int
		expect []int
		fn     func(elem int) bool
	}{
		{source: []int{1, 2, 3, 4, 5, 6}, expect: []int{1, 3, 5}, fn: func(elem int) bool {
			return elem%2 != 0
		}},
		{source: []int{1, 2, 3, 4, 5, 6}, expect: []int{2, 4, 6}, fn: func(elem int) bool {
			return elem%2 == 0
		}},
	}

	for _, test := range tests {
		var actual = slicex.Filter(test.source, test.fn)
		if !slicex.Equals(actual, test.expect, IntEqual) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}
