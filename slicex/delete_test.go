package slicex_test

import (
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestDelete(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		idx      int
		expected []int
	}{
		{
			name:     "删除第一个元素",
			source:   []int{1, 2, 3},
			idx:      0,
			expected: []int{2, 3},
		},
		{
			name:     "删除中间元素",
			source:   []int{1, 2, 3},
			idx:      1,
			expected: []int{1, 3},
		},
		{
			name:     "删除最后一个元素",
			source:   []int{1, 2, 3},
			idx:      2,
			expected: []int{1, 2},
		},
		{
			name:     "删除超出范围的索引",
			source:   []int{1, 2, 3},
			idx:      3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "删除负数索引",
			source:   []int{1, 2, 3},
			idx:      -1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "删除很大的索引",
			source:   []int{1, 2, 3},
			idx:      100,
			expected: []int{1, 2, 3},
		},
		{
			name:     "空切片删除",
			source:   []int{},
			idx:      0,
			expected: []int{},
		},
		{
			name:     "空切片删除负数索引",
			source:   []int{},
			idx:      -1,
			expected: []int{},
		},
		{
			name:     "单元素切片删除第一个",
			source:   []int{42},
			idx:      0,
			expected: []int{},
		},
		{
			name:     "单元素切片删除超出范围",
			source:   []int{42},
			idx:      1,
			expected: []int{42},
		},
		{
			name:     "两个元素切片删除第一个",
			source:   []int{1, 2},
			idx:      0,
			expected: []int{2},
		},
		{
			name:     "两个元素切片删除第二个",
			source:   []int{1, 2},
			idx:      1,
			expected: []int{1},
		},
		{
			name:     "nil切片删除",
			source:   nil,
			idx:      0,
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Delete(test.source, test.idx)
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestDeleteWithStrings(t *testing.T) {
	// 测试字符串类型的Delete
	var tests = []struct {
		name     string
		source   []string
		idx      int
		expected []string
	}{
		{
			name:     "删除字符串元素",
			source:   []string{"hello", "world", "test"},
			idx:      1,
			expected: []string{"hello", "test"},
		},
		{
			name:     "删除空字符串",
			source:   []string{"hello", "", "world"},
			idx:      1,
			expected: []string{"hello", "world"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Delete(test.source, test.idx)
			if !slicex.Equals(actual, test.expected, StringEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}
