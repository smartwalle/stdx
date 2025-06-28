package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestDelete(t *testing.T) {
	var tests = []struct {
		source []int
		idx    int
		expect []int
	}{
		{source: []int{1, 2, 3}, idx: 0, expect: []int{2, 3}},
		{source: []int{1, 2, 3}, idx: 1, expect: []int{1, 3}},
		{source: []int{1, 2, 3}, idx: 2, expect: []int{1, 2}},
		{source: []int{1, 2, 3}, idx: 3, expect: []int{1, 2, 3}},
		{source: []int{1, 2, 3}, idx: -1, expect: []int{1, 2, 3}},
		{source: []int{}, idx: 0, expect: []int{}},
		{source: []int{}, idx: 1, expect: []int{}},
	}

	for _, test := range tests {
		var actual = slicex.Delete(test.source, test.idx)
		if !slicex.Equals(actual, test.expect, IntEqual) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}
