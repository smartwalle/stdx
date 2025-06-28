package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestAppend(t *testing.T) {
	var tests = []struct {
		source []int
		append []int
		expect []int
	}{
		{source: []int{1, 2, 3}, append: []int{4, 5, 6}, expect: []int{1, 2, 3, 4, 5, 6}},
		{source: []int{}, append: []int{4, 5, 6}, expect: []int{4, 5, 6}},
	}

	for _, test := range tests {
		var actual = slicex.Append(test.source, test.append...)
		if !slicex.Equals(actual, test.expect, IntEqual) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}

func BenchmarkStdAppend(b *testing.B) {
	var ints []int
	for i := 0; i < b.N; i++ {
		ints = append(ints, i)
	}
}

func BenchmarkAppend(b *testing.B) {
	var ints []int
	for i := 0; i < b.N; i++ {
		ints = slicex.Append(ints, i)
	}
}
