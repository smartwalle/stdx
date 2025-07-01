package slicex_test

import (
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestAppend(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		append   []int
		expected []int
	}{
		{
			name:     "正常追加",
			source:   []int{1, 2, 3},
			append:   []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "空源切片追加",
			source:   []int{},
			append:   []int{4, 5, 6},
			expected: []int{4, 5, 6},
		},
		{
			name:     "追加空切片",
			source:   []int{1, 2, 3},
			append:   []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "空源切片追加空切片",
			source:   []int{},
			append:   []int{},
			expected: []int{},
		},
		{
			name:     "nil源切片追加",
			source:   nil,
			append:   []int{4, 5, 6},
			expected: []int{4, 5, 6},
		},
		{
			name:     "追加nil切片",
			source:   []int{1, 2, 3},
			append:   nil,
			expected: []int{1, 2, 3},
		},
		{
			name:     "nil源切片追加nil切片",
			source:   nil,
			append:   nil,
			expected: []int{},
		},
		{
			name:     "单元素追加",
			source:   []int{1, 2, 3},
			append:   []int{4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "追加单元素",
			source:   []int{1},
			append:   []int{2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "单元素追加单元素",
			source:   []int{1},
			append:   []int{2},
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Append(test.source, test.append...)
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestAppendWithStrings(t *testing.T) {
	// 测试字符串类型的Append
	var tests = []struct {
		name     string
		source   []string
		append   []string
		expected []string
	}{
		{
			name:     "字符串追加",
			source:   []string{"a", "b", "c"},
			append:   []string{"d", "e", "f"},
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name:     "空字符串切片追加",
			source:   []string{},
			append:   []string{"hello", "world"},
			expected: []string{"hello", "world"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Append(test.source, test.append...)
			if !slicex.Equals(actual, test.expected, StringEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
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
