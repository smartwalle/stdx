package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestPick(t *testing.T) {
	var values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var tests = []struct {
		offset int
		limit  int
		expect []int
	}{
		{offset: 0, limit: 0, expect: []int{}},
		{offset: 1, limit: 0, expect: []int{}},
		{offset: 0, limit: 1, expect: []int{1}},
		{offset: 1, limit: 1, expect: []int{2}},
		{offset: 1, limit: 2, expect: []int{2, 3}},
		{offset: 8, limit: 2, expect: []int{9}},
		{offset: 9, limit: 1, expect: []int{}},
		{offset: 5, limit: 2, expect: []int{6, 7}},
		{offset: 1, limit: 3, expect: []int{2, 3, 4}},
		{offset: 3, limit: 3, expect: []int{4, 5, 6}},
		{offset: 3, limit: 4, expect: []int{4, 5, 6, 7}},
		{offset: 4, limit: 4, expect: []int{5, 6, 7, 8}},
		{offset: 4, limit: 5, expect: []int{5, 6, 7, 8, 9}},
		{offset: 4, limit: 6, expect: []int{5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		var actual = slicex.Pick(values, test.offset, test.limit)
		if !slicex.Equals(actual, test.expect, IntEqual) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}
