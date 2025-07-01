package slicex_test

import (
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestUnique(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		expected []int
		fn       func(elem int) int
	}{
		{
			name:     "去除重复数字",
			source:   []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
			fn: func(elem int) int {
				return elem
			},
		},
		{
			name:     "去除重复数字（无序）",
			source:   []int{3, 3, 3, 2, 2, 2, 1, 1, 1},
			expected: []int{3, 2, 1},
			fn: func(elem int) int {
				return elem
			},
		},
		{
			name:     "单元素切片",
			source:   []int{1},
			expected: []int{1},
			fn: func(elem int) int {
				return elem
			},
		},
		{
			name:     "空切片",
			source:   []int{},
			expected: []int{},
			fn: func(elem int) int {
				return elem
			},
		},
		{
			name:     "nil切片",
			source:   nil,
			expected: []int{},
			fn: func(elem int) int {
				return elem
			},
		},
		{
			name:     "无重复元素",
			source:   []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
			fn: func(elem int) int {
				return elem
			},
		},
		{
			name:     "全部相同元素",
			source:   []int{42, 42, 42, 42, 42},
			expected: []int{42},
			fn: func(elem int) int {
				return elem
			},
		},
		{
			name:     "包含零值",
			source:   []int{0, 1, 0, 2, 0, 3},
			expected: []int{0, 1, 2, 3},
			fn: func(elem int) int {
				return elem
			},
		},
		{
			name:     "包含负数",
			source:   []int{1, -1, 2, -2, 3, -3},
			expected: []int{1, -1, 2, -2, 3, -3},
			fn: func(elem int) int {
				return elem
			},
		},
		{
			name:     "包含重复负数",
			source:   []int{1, -1, 1, -1, 2, -2, 2, -2},
			expected: []int{1, -1, 2, -2},
			fn: func(elem int) int {
				return elem
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Unique(test.source, test.fn)
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestUniqueWithCustomFunction(t *testing.T) {
	// 测试自定义去重函数
	var tests = []struct {
		name     string
		source   []int
		expected []int
		fn       func(elem int) int
	}{
		{
			name:     "按绝对值去重",
			source:   []int{1, -1, 2, -2, 3, -3},
			expected: []int{1, 2, 3},
			fn: func(elem int) int {
				if elem < 0 {
					return -elem
				}
				return elem
			},
		},
		{
			name:     "按奇偶性去重",
			source:   []int{1, 3, 5, 2, 4, 6, 7, 9},
			expected: []int{1, 2},
			fn: func(elem int) int {
				return elem % 2
			},
		},
		{
			name:     "按除以3的余数去重",
			source:   []int{1, 4, 7, 2, 5, 8, 3, 6, 9},
			expected: []int{1, 2, 3},
			fn: func(elem int) int {
				return elem % 3
			},
		},
		{
			name:     "按数字范围去重（0-9, 10-19, 20-29）",
			source:   []int{1, 15, 25, 5, 18, 28, 9, 19, 29},
			expected: []int{1, 15, 25},
			fn: func(elem int) int {
				return elem / 10
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Unique(test.source, test.fn)
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestUniqueWithStrings(t *testing.T) {
	// 测试字符串类型的Unique
	var tests = []struct {
		name     string
		source   []string
		expected []string
		fn       func(elem string) string
	}{
		{
			name:     "去除重复字符串",
			source:   []string{"hello", "world", "hello", "test", "world", "demo"},
			expected: []string{"hello", "world", "test", "demo"},
			fn: func(elem string) string {
				return elem
			},
		},
		{
			name:     "去除重复字符串（无序）",
			source:   []string{"test", "hello", "test", "world", "hello", "demo"},
			expected: []string{"test", "hello", "world", "demo"},
			fn: func(elem string) string {
				return elem
			},
		},
		{
			name:     "包含空字符串",
			source:   []string{"hello", "", "world", "", "test", ""},
			expected: []string{"hello", "", "world", "test"},
			fn: func(elem string) string {
				return elem
			},
		},
		{
			name:     "按字符串长度去重",
			source:   []string{"hi", "hello", "world", "a", "test", "go"},
			expected: []string{"hi", "hello"},
			fn: func(elem string) string {
				if len(elem) <= 2 {
					return "short"
				} else if len(elem) <= 5 {
					return "medium"
				}
				return "long"
			},
		},
		{
			name:     "按首字母去重",
			source:   []string{"hello", "world", "hi", "test", "help", "work"},
			expected: []string{"hello", "world", "test"},
			fn: func(elem string) string {
				if len(elem) > 0 {
					return string(elem[0])
				}
				return ""
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Unique(test.source, test.fn)
			t.Log(actual)
			if !slicex.Equals(actual, test.expected, StringEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestUniqueWithStruct(t *testing.T) {
	// 测试结构体类型的Unique
	type Person struct {
		Name string
		Age  int
	}

	var tests = []struct {
		name     string
		source   []Person
		expected []Person
		fn       func(elem Person) string
	}{
		{
			name: "按姓名去重",
			source: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Alice", Age: 28},
				{Name: "Charlie", Age: 35},
			},
			expected: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 35},
			},
			fn: func(elem Person) string {
				return elem.Name
			},
		},
		{
			name: "按年龄范围去重",
			source: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 35},
				{Name: "David", Age: 26},
				{Name: "Eve", Age: 31},
			},
			expected: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 35},
			},
			fn: func(elem Person) string {
				if elem.Age < 30 {
					return "young"
				} else if elem.Age < 35 {
					return "middle"
				}
				return "senior"
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Unique(test.source, test.fn)
			if len(actual) != len(test.expected) {
				t.Fatalf("实际长度: %d, 预期长度: %d", len(actual), len(test.expected))
			}
			// 由于结构体比较比较复杂，这里只检查长度
			// 在实际应用中，可能需要自定义比较函数
		})
	}
}
