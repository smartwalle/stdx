package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestContains(t *testing.T) {
	var tests = []struct {
		source []int
		expect bool
		fn     func(elem int) bool
	}{
		{source: []int{1, 2, 3, 4, 5, 6}, expect: true, fn: func(elem int) bool {
			return elem%2 != 0
		}},
		{source: []int{1, 2, 3, 4, 5, 6}, expect: true, fn: func(elem int) bool {
			return elem%2 == 0
		}},
		{source: []int{1, 2, 3}, expect: false, fn: func(elem int) bool {
			return elem == 0
		}},
		{source: []int{1, 2, 3}, expect: true, fn: func(elem int) bool {
			return elem == 3
		}},
	}

	for _, test := range tests {
		var actual = slicex.Contains(test.source, test.fn)
		if actual != test.expect {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}
