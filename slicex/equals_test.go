package slicex_test

import (
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestEquals(t *testing.T) {
	var tests = []struct {
		name     string
		s1       []int
		s2       []int
		fn       func(elem1, elem2 int) bool
		expected bool
	}{
		{
			name:     "相同长度的相等切片",
			s1:       []int{1, 2, 3},
			s2:       []int{1, 2, 3},
			fn:       IntEqual,
			expected: true,
		},
		{
			name:     "相同长度的不相等切片",
			s1:       []int{1, 2, 3},
			s2:       []int{1, 2, 4},
			fn:       IntEqual,
			expected: false,
		},
		{
			name:     "不同长度的切片",
			s1:       []int{1, 2, 3},
			s2:       []int{1, 2},
			fn:       IntEqual,
			expected: false,
		},
		{
			name:     "空切片与空切片",
			s1:       []int{},
			s2:       []int{},
			fn:       IntEqual,
			expected: true,
		},
		{
			name:     "空切片与非空切片",
			s1:       []int{},
			s2:       []int{1, 2, 3},
			fn:       IntEqual,
			expected: false,
		},
		{
			name:     "nil切片与nil切片",
			s1:       nil,
			s2:       nil,
			fn:       IntEqual,
			expected: true,
		},
		{
			name:     "nil切片与空切片",
			s1:       nil,
			s2:       []int{},
			fn:       IntEqual,
			expected: true,
		},
		{
			name:     "nil切片与非空切片",
			s1:       nil,
			s2:       []int{1, 2, 3},
			fn:       IntEqual,
			expected: false,
		},
		{
			name:     "单元素切片相等",
			s1:       []int{42},
			s2:       []int{42},
			fn:       IntEqual,
			expected: true,
		},
		{
			name:     "单元素切片不相等",
			s1:       []int{42},
			s2:       []int{43},
			fn:       IntEqual,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := slicex.Equals(test.s1, test.s2, test.fn)
			if actual != test.expected {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestEqualsWithCustomFunction(t *testing.T) {
	// 测试自定义比较函数
	var tests = []struct {
		name     string
		s1       []string
		s2       []string
		fn       func(elem1, elem2 string) bool
		expected bool
	}{
		{
			name: "忽略大小写比较",
			s1:   []string{"Hello", "World"},
			s2:   []string{"hello", "world"},
			fn: func(elem1, elem2 string) bool {
				return elem1 == elem2
			},
			expected: false,
		},
		{
			name: "忽略大小写比较（相等）",
			s1:   []string{"Hello", "World"},
			s2:   []string{"Hello", "World"},
			fn: func(elem1, elem2 string) bool {
				return elem1 == elem2
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := slicex.Equals(test.s1, test.s2, test.fn)
			if actual != test.expected {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}
