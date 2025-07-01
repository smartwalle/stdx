package slicex_test

import (
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestInsert(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		insert   []int
		idx      int
		expected []int
	}{
		{
			name:     "空切片插入到负数位置",
			source:   []int{},
			insert:   []int{4, 5, 6},
			idx:      -2,
			expected: []int{4, 5, 6},
		},
		{
			name:     "空切片插入到-1位置",
			source:   []int{},
			insert:   []int{4, 5, 6},
			idx:      -1,
			expected: []int{4, 5, 6},
		},
		{
			name:     "空切片插入到0位置",
			source:   []int{},
			insert:   []int{4, 5, 6},
			idx:      0,
			expected: []int{4, 5, 6},
		},
		{
			name:     "空切片插入到正数位置",
			source:   []int{},
			insert:   []int{4, 5, 6},
			idx:      1,
			expected: []int{4, 5, 6},
		},
		{
			name:     "插入空切片到负数位置",
			source:   []int{1, 2, 3},
			insert:   []int{},
			idx:      -2,
			expected: []int{1, 2, 3},
		},
		{
			name:     "插入空切片到-1位置",
			source:   []int{1, 2, 3},
			insert:   []int{},
			idx:      -1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "插入空切片到0位置",
			source:   []int{1, 2, 3},
			insert:   []int{},
			idx:      0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "插入空切片到正数位置",
			source:   []int{1, 2, 3},
			insert:   []int{},
			idx:      1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "插入到负数位置（开头）",
			source:   []int{1, 2, 3},
			insert:   []int{4, 5, 6},
			idx:      -2,
			expected: []int{4, 5, 6, 1, 2, 3},
		},
		{
			name:     "插入到-1位置（开头）",
			source:   []int{1, 2, 3},
			insert:   []int{4, 5, 6},
			idx:      -1,
			expected: []int{4, 5, 6, 1, 2, 3},
		},
		{
			name:     "插入到0位置（开头）",
			source:   []int{1, 2, 3},
			insert:   []int{4, 5, 6},
			idx:      0,
			expected: []int{4, 5, 6, 1, 2, 3},
		},
		{
			name:     "插入到1位置（中间）",
			source:   []int{1, 2, 3},
			insert:   []int{4, 5, 6},
			idx:      1,
			expected: []int{1, 4, 5, 6, 2, 3},
		},
		{
			name:     "插入到2位置（中间）",
			source:   []int{1, 2, 3},
			insert:   []int{4, 5, 6},
			idx:      2,
			expected: []int{1, 2, 4, 5, 6, 3},
		},
		{
			name:     "插入到3位置（结尾）",
			source:   []int{1, 2, 3},
			insert:   []int{4, 5, 6},
			idx:      3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "插入到超出范围位置（结尾）",
			source:   []int{1, 2, 3},
			insert:   []int{4, 5, 6},
			idx:      4,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "插入到很大位置（结尾）",
			source:   []int{1, 2, 3},
			insert:   []int{4, 5, 6},
			idx:      5,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "插入到更大位置（结尾）",
			source:   []int{1, 2, 3},
			insert:   []int{4, 5, 6},
			idx:      6,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "单元素切片插入到开头",
			source:   []int{42},
			insert:   []int{1, 2, 3},
			idx:      0,
			expected: []int{1, 2, 3, 42},
		},
		{
			name:     "单元素切片插入到结尾",
			source:   []int{42},
			insert:   []int{1, 2, 3},
			idx:      1,
			expected: []int{42, 1, 2, 3},
		},
		{
			name:     "插入单元素到开头",
			source:   []int{1, 2, 3},
			insert:   []int{42},
			idx:      0,
			expected: []int{42, 1, 2, 3},
		},
		{
			name:     "插入单元素到中间",
			source:   []int{1, 2, 3},
			insert:   []int{42},
			idx:      1,
			expected: []int{1, 42, 2, 3},
		},
		{
			name:     "插入单元素到结尾",
			source:   []int{1, 2, 3},
			insert:   []int{42},
			idx:      3,
			expected: []int{1, 2, 3, 42},
		},
		{
			name:     "nil切片插入",
			source:   nil,
			insert:   []int{1, 2, 3},
			idx:      0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "插入nil切片",
			source:   []int{1, 2, 3},
			insert:   nil,
			idx:      1,
			expected: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Insert(test.source, test.idx, test.insert...)
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestInsertWithStrings(t *testing.T) {
	// 测试字符串类型的Insert
	var tests = []struct {
		name     string
		source   []string
		insert   []string
		idx      int
		expected []string
	}{
		{
			name:     "字符串插入到开头",
			source:   []string{"world", "test"},
			insert:   []string{"hello"},
			idx:      0,
			expected: []string{"hello", "world", "test"},
		},
		{
			name:     "字符串插入到中间",
			source:   []string{"hello", "test"},
			insert:   []string{"world"},
			idx:      1,
			expected: []string{"hello", "world", "test"},
		},
		{
			name:     "插入空字符串",
			source:   []string{"hello", "world"},
			insert:   []string{""},
			idx:      1,
			expected: []string{"hello", "", "world"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Insert(test.source, test.idx, test.insert...)
			if !slicex.Equals(actual, test.expected, StringEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}
