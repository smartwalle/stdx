package slicex_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestMap(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		expected []string
		fn       func(elem int) string
	}{
		{
			name:     "数字转字符串（+1）",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: []string{"2", "3", "4", "5", "6", "7"},
			fn: func(elem int) string {
				return strconv.FormatInt(int64(elem+1), 10)
			},
		},
		{
			name:     "数字转字符串（*2）",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: []string{"2", "4", "6", "8", "10", "12"},
			fn: func(elem int) string {
				return strconv.FormatInt(int64(elem*2), 10)
			},
		},
		{
			name:     "空切片映射",
			source:   []int{},
			expected: []string{},
			fn: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "nil切片映射",
			source:   nil,
			expected: []string{},
			fn: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "单元素切片映射",
			source:   []int{42},
			expected: []string{"42"},
			fn: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "负数映射",
			source:   []int{-1, -2, -3},
			expected: []string{"-1", "-2", "-3"},
			fn: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "零值映射",
			source:   []int{0, 1, 0, 2},
			expected: []string{"0", "1", "0", "2"},
			fn: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "条件映射（偶数转even，奇数转odd）",
			source:   []int{1, 2, 3, 4, 5},
			expected: []string{"odd", "even", "odd", "even", "odd"},
			fn: func(elem int) string {
				if elem%2 == 0 {
					return "even"
				}
				return "odd"
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Map(test.source, test.fn)
			if !slicex.Equals(actual, test.expected, StringEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestMapMatched(t *testing.T) {
	var tests = []struct {
		name       string
		source     []int
		expected   []string
		filterFunc func(elem int) bool
		mapFunc    func(elem int) string
	}{
		{
			name:     "过滤所有元素并映射",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: []string{"1", "2", "3", "4", "5", "6"},
			filterFunc: func(elem int) bool {
				return true
			},
			mapFunc: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "所有元素都不匹配",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: []string{},
			filterFunc: func(elem int) bool {
				return false
			},
			mapFunc: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "过滤所有元素并映射（*2）",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: []string{"2", "4", "6", "8", "10", "12"},
			filterFunc: func(elem int) bool {
				return true
			},
			mapFunc: func(elem int) string {
				return strconv.FormatInt(int64(elem*2), 10)
			},
		},
		{
			name:     "过滤偶数并映射",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: []string{"2", "4", "6"},
			filterFunc: func(elem int) bool {
				return elem%2 == 0
			},
			mapFunc: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "过滤奇数并映射",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: []string{"1", "3", "5"},
			filterFunc: func(elem int) bool {
				return elem%2 != 0
			},
			mapFunc: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "过滤大于3的数并映射",
			source:   []int{1, 2, 3, 4, 5, 6},
			expected: []string{"4", "5", "6"},
			filterFunc: func(elem int) bool {
				return elem > 3
			},
			mapFunc: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "空切片过滤映射",
			source:   []int{},
			expected: []string{},
			filterFunc: func(elem int) bool {
				return elem > 0
			},
			mapFunc: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "nil切片过滤映射",
			source:   nil,
			expected: []string{},
			filterFunc: func(elem int) bool {
				return elem > 0
			},
			mapFunc: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "过滤所有元素（条件为false）",
			source:   []int{1, 2, 3, 4, 5},
			expected: []string{},
			filterFunc: func(elem int) bool {
				return false
			},
			mapFunc: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
		{
			name:     "过滤负数并映射",
			source:   []int{1, -2, 3, -4, 5, -6},
			expected: []string{"-2", "-4", "-6"},
			filterFunc: func(elem int) bool {
				return elem < 0
			},
			mapFunc: func(elem int) string {
				return strconv.FormatInt(int64(elem), 10)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.MapMatched(test.source, test.filterFunc, test.mapFunc)
			if !slicex.Equals(actual, test.expected, StringEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestMapWithIntToInt(t *testing.T) {
	// 测试int到int的映射
	var tests = []struct {
		name     string
		source   []int
		expected []int
		fn       func(elem int) int
	}{
		{
			name:     "数字平方",
			source:   []int{1, 2, 3, 4, 5},
			expected: []int{1, 4, 9, 16, 25},
			fn: func(elem int) int {
				return elem * elem
			},
		},
		{
			name:     "数字加1",
			source:   []int{1, 2, 3, 4, 5},
			expected: []int{2, 3, 4, 5, 6},
			fn: func(elem int) int {
				return elem + 1
			},
		},
		{
			name:     "空切片",
			source:   []int{},
			expected: []int{},
			fn: func(elem int) int {
				return elem * 2
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Map(test.source, test.fn)
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestMapWithStructs(t *testing.T) {
	// 测试自定义结构体类型的Map
	type Person struct {
		Name string
		Age  int
	}

	var tests = []struct {
		name     string
		source   []Person
		expected []string
		fn       func(elem Person) string
	}{
		{
			name: "提取姓名",
			source: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 35},
			},
			expected: []string{"Alice", "Bob", "Charlie"},
			fn: func(elem Person) string {
				return elem.Name
			},
		},
		{
			name: "格式化姓名和年龄",
			source: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 35},
			},
			expected: []string{"Alice(25)", "Bob(30)", "Charlie(35)"},
			fn: func(elem Person) string {
				return fmt.Sprintf("%s(%d)", elem.Name, elem.Age)
			},
		},
		{
			name:     "空结构体切片",
			source:   []Person{},
			expected: []string{},
			fn: func(elem Person) string {
				return elem.Name
			},
		},
		{
			name: "格式化姓名和年龄",
			source: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 17},
				{Name: "Charlie", Age: 30},
				{Name: "David", Age: 16},
			},
			expected: []string{"Alice(25)", "Bob(17)", "Charlie(30)", "David(16)"},
			fn: func(elem Person) string {
				return fmt.Sprintf("%s(%d)", elem.Name, elem.Age)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.Map(test.source, test.fn)
			if !slicex.Equals(actual, test.expected, StringEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestMapMatchedWithStructs(t *testing.T) {
	// 测试自定义结构体类型的MapMatched
	type Person struct {
		Name string
		Age  int
	}

	var tests = []struct {
		name      string
		source    []Person
		expected  []string
		predicate func(elem Person) bool
		mapFunc   func(elem Person) string
	}{
		{
			name: "过滤成年人并提取姓名",
			source: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 17},
				{Name: "Charlie", Age: 30},
				{Name: "David", Age: 16},
			},
			expected: []string{"Alice", "Charlie"},
			predicate: func(elem Person) bool {
				return elem.Age >= 18
			},
			mapFunc: func(elem Person) string {
				return elem.Name
			},
		},
		{
			name: "过滤名字以A开头的并格式化",
			source: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 17},
				{Name: "Charlie", Age: 30},
				{Name: "David", Age: 16},
			},
			expected: []string{"Alice(25)"},
			predicate: func(elem Person) bool {
				return len(elem.Name) > 0 && elem.Name[0] == 'A'
			},
			mapFunc: func(elem Person) string {
				return fmt.Sprintf("%s(%d)", elem.Name, elem.Age)
			},
		},
		{
			name:     "空结构体切片",
			source:   []Person{},
			expected: []string{},
			predicate: func(elem Person) bool {
				return elem.Age >= 18
			},
			mapFunc: func(elem Person) string {
				return elem.Name
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicex.MapMatched(test.source, test.predicate, test.mapFunc)
			if !slicex.Equals(actual, test.expected, StringEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}
