package slicex_test

import (
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestContains(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		expected bool
		fn       func(elem int) bool
	}{
		{
			name:     "包含奇数",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: true,
			fn: func(elem int) bool {
				return elem%2 != 0
			},
		},
		{
			name:     "包含偶数",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: true,
			fn: func(elem int) bool {
				return elem%2 == 0
			},
		},
		{
			name:     "不包含零",
			source:   []int{1, 2, 3},
			expected: false,
			fn: func(elem int) bool {
				return elem == 0
			},
		},
		{
			name:     "包含特定值",
			source:   []int{1, 2, 3},
			expected: true,
			fn: func(elem int) bool {
				return elem == 3
			},
		},
		{
			name:     "空切片",
			source:   []int{},
			expected: false,
			fn: func(elem int) bool {
				return elem == 1
			},
		},
		{
			name:     "nil切片",
			source:   nil,
			expected: false,
			fn: func(elem int) bool {
				return elem == 1
			},
		},
		{
			name:     "包含大于5的数",
			source:   []int{1, 2, 3, 4, 5, 6, 7},
			expected: true,
			fn: func(elem int) bool {
				return elem > 5
			},
		},
		{
			name:     "不包含大于10的数",
			source:   []int{1, 2, 3, 4, 5, 6, 7},
			expected: false,
			fn: func(elem int) bool {
				return elem > 10
			},
		},
		{
			name:     "包含负数",
			source:   []int{1, -2, 3, -4, 5},
			expected: true,
			fn: func(elem int) bool {
				return elem < 0
			},
		},
		{
			name:     "单元素切片包含",
			source:   []int{42},
			expected: true,
			fn: func(elem int) bool {
				return elem == 42
			},
		},
		{
			name:     "单元素切片不包含",
			source:   []int{42},
			expected: false,
			fn: func(elem int) bool {
				return elem == 43
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Contains(test.source, test.fn)
			if actual != test.expected {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestContainsWithStrings(t *testing.T) {
	// 测试字符串类型的Contains
	var tests = []struct {
		name     string
		source   []string
		expected bool
		fn       func(elem string) bool
	}{
		{
			name:     "包含空字符串",
			source:   []string{"hello", "", "world"},
			expected: true,
			fn: func(elem string) bool {
				return elem == ""
			},
		},
		{
			name:     "包含以h开头的字符串",
			source:   []string{"hello", "world", "hi"},
			expected: true,
			fn: func(elem string) bool {
				return len(elem) > 0 && elem[0] == 'h'
			},
		},
		{
			name:     "不包含长度大于10的字符串",
			source:   []string{"hello", "world", "hi"},
			expected: false,
			fn: func(elem string) bool {
				return len(elem) > 10
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Contains(test.source, test.fn)
			if actual != test.expected {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}
