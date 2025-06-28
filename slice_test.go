package stdx_test

import (
	"github.com/smartwalle/stdx"
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
		var actual = stdx.Append(test.source, test.append...)
		if !Equal(actual, test.expect) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}

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
		var actual = stdx.Split(values, test.size)[test.index]
		if !Equal(actual, test.expect) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}

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
		var actual = stdx.Pick(values, test.offset, test.limit)
		if !Equal(actual, test.expect) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}

func Equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
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
		ints = stdx.Append(ints, i)
	}
}
