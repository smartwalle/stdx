package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestCompare(t *testing.T) {
	var tests = []struct {
		source []int
		expect int
		fn     func(elem1, elem2 int) bool
	}{
		{source: []int{1, 2, 3, 4, 5, 6}, expect: 6, fn: func(elem1, elem2 int) bool {
			return elem1 > elem2
		}},
		{source: []int{1, 2, 3, 4, 5, 6}, expect: 1, fn: func(elem1, elem2 int) bool {
			return elem1 < elem2
		}},
		{source: []int{1}, expect: 1, fn: func(elem1, elem2 int) bool {
			return elem1 < elem2
		}},
		{source: []int{6}, expect: 6, fn: func(elem1, elem2 int) bool {
			return elem1 < elem2
		}},
		{source: []int{}, expect: 0, fn: func(elem1, elem2 int) bool {
			return elem1 < elem2
		}},
	}

	for _, test := range tests {
		var actual = slicex.Compare(test.source, test.fn)
		if actual != test.expect {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		source []int
		expect int
	}{
		{source: []int{1, 2, 3, 4, 5, 6}, expect: 1},
		{source: []int{9, 8, 4, 2, 6}, expect: 2},
		{source: []int{1}, expect: 1},
		{source: []int{6}, expect: 6},
		{source: []int{}, expect: 0},
	}

	for _, test := range tests {
		var actual = slicex.Min(test.source)
		if actual != test.expect {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}

func TestMax(t *testing.T) {
	var tests = []struct {
		source []int
		expect int
	}{
		{source: []int{1, 2, 3, 4, 5, 6}, expect: 6},
		{source: []int{9, 8, 4, 2, 6}, expect: 9},
		{source: []int{1}, expect: 1},
		{source: []int{6}, expect: 6},
		{source: []int{}, expect: 0},
	}

	for _, test := range tests {
		var actual = slicex.Max(test.source)
		if actual != test.expect {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}
