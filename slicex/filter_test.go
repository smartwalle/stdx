package slicex_test

import (
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestFilter(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		expected []int
		fn       func(elem int) bool
	}{
		{
			name:     "过滤奇数",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: []int{1, 3, 5},
			fn: func(elem int) bool {
				return elem%2 != 0
			},
		},
		{
			name:     "过滤偶数",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 4, 6},
			fn: func(elem int) bool {
				return elem%2 == 0
			},
		},
		{
			name:     "过滤大于3的数",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: []int{4, 5, 6},
			fn: func(elem int) bool {
				return elem > 3
			},
		},
		{
			name:     "过滤等于特定值的数",
			source:   []int{1, 2, 3, 2, 4, 2},
			expected: []int{2, 2, 2},
			fn: func(elem int) bool {
				return elem == 2
			},
		},
		{
			name:     "空切片过滤",
			source:   []int{},
			expected: []int{},
			fn: func(elem int) bool {
				return elem > 0
			},
		},
		{
			name:     "nil切片过滤",
			source:   nil,
			expected: []int{},
			fn: func(elem int) bool {
				return elem > 0
			},
		},
		{
			name:     "过滤负数",
			source:   []int{1, -2, 3, -4, 5, -6},
			expected: []int{-2, -4, -6},
			fn: func(elem int) bool {
				return elem < 0
			},
		},
		{
			name:     "过滤所有元素（条件为true）",
			source:   []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
			fn: func(elem int) bool {
				return true
			},
		},
		{
			name:     "过滤所有元素（条件为false）",
			source:   []int{1, 2, 3, 4, 5},
			expected: []int{},
			fn: func(elem int) bool {
				return false
			},
		},
		{
			name:     "单元素切片过滤",
			source:   []int{42},
			expected: []int{42},
			fn: func(elem int) bool {
				return elem == 42
			},
		},
		{
			name:     "单元素切片不匹配",
			source:   []int{42},
			expected: []int{},
			fn: func(elem int) bool {
				return elem == 43
			},
		},
		{
			name:     "过滤零值",
			source:   []int{0, 1, 0, 2, 0, 3},
			expected: []int{0, 0, 0},
			fn: func(elem int) bool {
				return elem == 0
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Filter(test.source, test.fn)
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestFilterWithStrings(t *testing.T) {
	// 测试字符串类型的Filter
	var tests = []struct {
		name     string
		source   []string
		expected []string
		fn       func(elem string) bool
	}{
		{
			name:     "过滤空字符串",
			source:   []string{"hello", "", "world", "", "test"},
			expected: []string{"", ""},
			fn: func(elem string) bool {
				return elem == ""
			},
		},
		{
			name:     "过滤非空字符串",
			source:   []string{"hello", "", "world", "", "test"},
			expected: []string{"hello", "world", "test"},
			fn: func(elem string) bool {
				return elem != ""
			},
		},
		{
			name:     "过滤以h开头的字符串",
			source:   []string{"hello", "world", "hi", "test"},
			expected: []string{"hello", "hi"},
			fn: func(elem string) bool {
				return len(elem) > 0 && elem[0] == 'h'
			},
		},
		{
			name:     "过滤长度大于3的字符串",
			source:   []string{"hi", "hello", "world", "a", "test"},
			expected: []string{"hello", "world", "test"},
			fn: func(elem string) bool {
				return len(elem) > 3
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Filter(test.source, test.fn)
			if !slicex.Equals(actual, test.expected, StringEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}
