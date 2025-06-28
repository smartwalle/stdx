package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestInsert(t *testing.T) {
	var tests = []struct {
		source []int
		insert []int
		idx    int
		expect []int
	}{
		{source: []int{}, insert: []int{4, 5, 6}, idx: -2, expect: []int{4, 5, 6}},
		{source: []int{}, insert: []int{4, 5, 6}, idx: -1, expect: []int{4, 5, 6}},
		{source: []int{}, insert: []int{4, 5, 6}, idx: 0, expect: []int{4, 5, 6}},
		{source: []int{}, insert: []int{4, 5, 6}, idx: 1, expect: []int{4, 5, 6}},
		{source: []int{}, insert: []int{4, 5, 6}, idx: 2, expect: []int{4, 5, 6}},
		{source: []int{1, 2, 3}, insert: []int{}, idx: -2, expect: []int{1, 2, 3}},
		{source: []int{1, 2, 3}, insert: []int{}, idx: -1, expect: []int{1, 2, 3}},
		{source: []int{1, 2, 3}, insert: []int{}, idx: 0, expect: []int{1, 2, 3}},
		{source: []int{1, 2, 3}, insert: []int{}, idx: 1, expect: []int{1, 2, 3}},
		{source: []int{1, 2, 3}, insert: []int{}, idx: 2, expect: []int{1, 2, 3}},
		{source: []int{1, 2, 3}, insert: []int{4, 5, 6}, idx: -2, expect: []int{4, 5, 6, 1, 2, 3}},
		{source: []int{1, 2, 3}, insert: []int{4, 5, 6}, idx: -1, expect: []int{4, 5, 6, 1, 2, 3}},
		{source: []int{1, 2, 3}, insert: []int{4, 5, 6}, idx: 0, expect: []int{4, 5, 6, 1, 2, 3}},
		{source: []int{1, 2, 3}, insert: []int{4, 5, 6}, idx: 1, expect: []int{1, 4, 5, 6, 2, 3}},
		{source: []int{1, 2, 3}, insert: []int{4, 5, 6}, idx: 2, expect: []int{1, 2, 4, 5, 6, 3}},
		{source: []int{1, 2, 3}, insert: []int{4, 5, 6}, idx: 3, expect: []int{1, 2, 3, 4, 5, 6}},
		{source: []int{1, 2, 3}, insert: []int{4, 5, 6}, idx: 4, expect: []int{1, 2, 3, 4, 5, 6}},
		{source: []int{1, 2, 3}, insert: []int{4, 5, 6}, idx: 5, expect: []int{1, 2, 3, 4, 5, 6}},
		{source: []int{1, 2, 3}, insert: []int{4, 5, 6}, idx: 6, expect: []int{1, 2, 3, 4, 5, 6}},
	}

	for _, test := range tests {
		var actual = slicex.Insert(test.source, test.idx, test.insert...)
		if !slicex.Equals(actual, test.expect, IntEqual) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}
