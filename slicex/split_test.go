package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestSplit(t *testing.T) {
	var values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	var tests = []struct {
		size   int
		index  int
		expect []int
	}{
		{size: 3, index: 0, expect: []int{1, 2, 3}},
		{size: 3, index: 1, expect: []int{4, 5, 6}},
		{size: 3, index: 5, expect: []int{16}},
		{size: 4, index: 3, expect: []int{13, 14, 15, 16}},
	}

	for _, test := range tests {
		var actual = slicex.Split(values, test.size)[test.index]
		if !slicex.Equals(actual, test.expect, IntEqual) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}
