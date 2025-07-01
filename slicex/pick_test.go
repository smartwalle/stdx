package slicex_test

import (
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestPick(t *testing.T) {
	var values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var tests = []struct {
		name     string
		offset   int
		limit    int
		expected []int
	}{
		{
			name:     "offset=0, limit=0",
			offset:   0,
			limit:    0,
			expected: []int{},
		},
		{
			name:     "offset=1, limit=0",
			offset:   1,
			limit:    0,
			expected: []int{},
		},
		{
			name:     "offset=0, limit=1",
			offset:   0,
			limit:    1,
			expected: []int{1},
		},
		{
			name:     "offset=1, limit=1",
			offset:   1,
			limit:    1,
			expected: []int{2},
		},
		{
			name:     "offset=1, limit=2",
			offset:   1,
			limit:    2,
			expected: []int{2, 3},
		},
		{
			name:     "offset=8, limit=2（超出范围）",
			offset:   8,
			limit:    2,
			expected: []int{9},
		},
		{
			name:     "offset=9, limit=1（超出范围）",
			offset:   9,
			limit:    1,
			expected: []int{},
		},
		{
			name:     "offset=5, limit=2",
			offset:   5,
			limit:    2,
			expected: []int{6, 7},
		},
		{
			name:     "offset=1, limit=3",
			offset:   1,
			limit:    3,
			expected: []int{2, 3, 4},
		},
		{
			name:     "offset=3, limit=3",
			offset:   3,
			limit:    3,
			expected: []int{4, 5, 6},
		},
		{
			name:     "offset=3, limit=4",
			offset:   3,
			limit:    4,
			expected: []int{4, 5, 6, 7},
		},
		{
			name:     "offset=4, limit=4",
			offset:   4,
			limit:    4,
			expected: []int{5, 6, 7, 8},
		},
		{
			name:     "offset=4, limit=5",
			offset:   4,
			limit:    5,
			expected: []int{5, 6, 7, 8, 9},
		},
		{
			name:     "offset=4, limit=6（超出范围）",
			offset:   4,
			limit:    6,
			expected: []int{5, 6, 7, 8, 9},
		},
		{
			name:     "offset=0, limit=9（全部）",
			offset:   0,
			limit:    9,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "offset=0, limit=10（超出范围）",
			offset:   0,
			limit:    10,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "offset=10, limit=1（超出范围）",
			offset:   10,
			limit:    1,
			expected: []int{},
		},
		{
			name:     "offset=-1, limit=1（负数offset）",
			offset:   -1,
			limit:    1,
			expected: []int{},
		},
		{
			name:     "offset=0, limit=-1（负数limit）",
			offset:   0,
			limit:    -1,
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Pick(values, test.offset, test.limit)
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestPickWithEmptySlice(t *testing.T) {
	// 测试空切片的情况
	var tests = []struct {
		name     string
		offset   int
		limit    int
		expected []int
	}{
		{
			name:     "空切片，offset=0, limit=0",
			offset:   0,
			limit:    0,
			expected: []int{},
		},
		{
			name:     "空切片，offset=0, limit=1",
			offset:   0,
			limit:    1,
			expected: []int{},
		},
		{
			name:     "空切片，offset=1, limit=1",
			offset:   1,
			limit:    1,
			expected: []int{},
		},
		{
			name:     "空切片，offset=-1, limit=1",
			offset:   -1,
			limit:    1,
			expected: []int{},
		},
	}

	emptySlice := []int{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Pick(emptySlice, test.offset, test.limit)
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestPickWithNilSlice(t *testing.T) {
	// 测试nil切片的情况
	var tests = []struct {
		name     string
		offset   int
		limit    int
		expected []int
	}{
		{
			name:     "nil切片，offset=0, limit=0",
			offset:   0,
			limit:    0,
			expected: []int{},
		},
		{
			name:     "nil切片，offset=0, limit=1",
			offset:   0,
			limit:    1,
			expected: []int{},
		},
		{
			name:     "nil切片，offset=1, limit=1",
			offset:   1,
			limit:    1,
			expected: []int{},
		},
	}

	var nilSlice []int
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Pick(nilSlice, test.offset, test.limit)
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestPickWithSingleElement(t *testing.T) {
	// 测试单元素切片
	singleSlice := []int{42}
	var tests = []struct {
		name     string
		offset   int
		limit    int
		expected []int
	}{
		{
			name:     "单元素，offset=0, limit=0",
			offset:   0,
			limit:    0,
			expected: []int{},
		},
		{
			name:     "单元素，offset=0, limit=1",
			offset:   0,
			limit:    1,
			expected: []int{42},
		},
		{
			name:     "单元素，offset=0, limit=2",
			offset:   0,
			limit:    2,
			expected: []int{42},
		},
		{
			name:     "单元素，offset=1, limit=1",
			offset:   1,
			limit:    1,
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Pick(singleSlice, test.offset, test.limit)
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestPickWithStrings(t *testing.T) {
	// 测试字符串类型的Pick
	values := []string{"hello", "world", "test", "example"}
	var tests = []struct {
		name     string
		offset   int
		limit    int
		expected []string
	}{
		{
			name:     "字符串切片，offset=0, limit=2",
			offset:   0,
			limit:    2,
			expected: []string{"hello", "world"},
		},
		{
			name:     "字符串切片，offset=1, limit=2",
			offset:   1,
			limit:    2,
			expected: []string{"world", "test"},
		},
		{
			name:     "字符串切片，offset=2, limit=3（超出范围）",
			offset:   2,
			limit:    3,
			expected: []string{"test", "example"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Pick(values, test.offset, test.limit)
			if !slicex.Equals(actual, test.expected, StringEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}
