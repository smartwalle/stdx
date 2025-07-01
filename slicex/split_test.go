package slicex_test

import (
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestSplit(t *testing.T) {
	var values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	var tests = []struct {
		name     string
		size     int
		index    int
		expected []int
	}{
		{
			name:     "size=3, index=0",
			size:     3,
			index:    0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "size=3, index=1",
			size:     3,
			index:    1,
			expected: []int{4, 5, 6},
		},
		{
			name:     "size=3, index=5（最后一个不完整块）",
			size:     3,
			index:    5,
			expected: []int{16},
		},
		{
			name:     "size=4, index=3（最后一个完整块）",
			size:     4,
			index:    3,
			expected: []int{13, 14, 15, 16},
		},
		{
			name:     "size=2, index=0",
			size:     2,
			index:    0,
			expected: []int{1, 2},
		},
		{
			name:     "size=2, index=7（最后一个块）",
			size:     2,
			index:    7,
			expected: []int{15, 16},
		},
		{
			name:     "size=5, index=0",
			size:     5,
			index:    0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "size=5, index=3（最后一个不完整块）",
			size:     5,
			index:    3,
			expected: []int{16},
		},
		{
			name:     "size=8, index=0",
			size:     8,
			index:    0,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "size=8, index=1（最后一个块）",
			size:     8,
			index:    1,
			expected: []int{9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			name:     "size=16, index=0（整个切片）",
			size:     16,
			index:    0,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			name:     "size=1, index=0",
			size:     1,
			index:    0,
			expected: []int{1},
		},
		{
			name:     "size=1, index=15（最后一个）",
			size:     1,
			index:    15,
			expected: []int{16},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := slicex.Split(values, test.size)
			if test.index >= len(result) {
				t.Fatalf("索引 %d 超出结果切片长度 %d", test.index, len(result))
			}
			var actual = result[test.index]
			if !slicex.Equals(actual, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestSplitWithEmptySlice(t *testing.T) {
	// 测试空切片的情况
	var tests = []struct {
		name     string
		size     int
		expected int // 期望的结果切片数量
	}{
		{
			name:     "空切片，size=1",
			size:     1,
			expected: 0,
		},
		{
			name:     "空切片，size=5",
			size:     5,
			expected: 0,
		},
		{
			name:     "空切片，size=0",
			size:     0,
			expected: 0,
		},
	}

	emptySlice := []int{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := slicex.Split(emptySlice, test.size)
			if len(result) != test.expected {
				t.Fatalf("实际结果数量: %d, 预期: %d", len(result), test.expected)
			}
		})
	}
}

func TestSplitWithNilSlice(t *testing.T) {
	// 测试nil切片的情况
	var tests = []struct {
		name     string
		size     int
		expected int // 期望的结果切片数量
	}{
		{
			name:     "nil切片，size=1",
			size:     1,
			expected: 0,
		},
		{
			name:     "nil切片，size=5",
			size:     5,
			expected: 0,
		},
	}

	var nilSlice []int
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := slicex.Split(nilSlice, test.size)
			if len(result) != test.expected {
				t.Fatalf("实际结果数量: %d, 预期: %d", len(result), test.expected)
			}
		})
	}
}

func TestSplitWithSingleElement(t *testing.T) {
	// 测试单元素切片
	singleSlice := []int{42}
	var tests = []struct {
		name     string
		size     int
		expected [][]int
	}{
		{
			name:     "单元素，size=1",
			size:     1,
			expected: [][]int{{42}},
		},
		{
			name:     "单元素，size=2",
			size:     2,
			expected: [][]int{{42}},
		},
		{
			name:     "单元素，size=5",
			size:     5,
			expected: [][]int{{42}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := slicex.Split(singleSlice, test.size)
			if len(result) != len(test.expected) {
				t.Fatalf("实际结果数量: %d, 预期: %d", len(result), len(test.expected))
			}
			for i, expected := range test.expected {
				if !slicex.Equals(result[i], expected, IntEqual) {
					t.Fatalf("索引 %d: 实际: %+v, 预期: %+v", i, result[i], expected)
				}
			}
		})
	}
}

func TestSplitWithSmallSlice(t *testing.T) {
	// 测试小切片的情况
	smallSlice := []int{1, 2, 3}
	var tests = []struct {
		name     string
		size     int
		expected [][]int
	}{
		{
			name:     "小切片，size=1",
			size:     1,
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "小切片，size=2",
			size:     2,
			expected: [][]int{{1, 2}, {3}},
		},
		{
			name:     "小切片，size=3",
			size:     3,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name:     "小切片，size=4",
			size:     4,
			expected: [][]int{{1, 2, 3}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := slicex.Split(smallSlice, test.size)
			if len(result) != len(test.expected) {
				t.Fatalf("实际结果数量: %d, 预期: %d", len(result), len(test.expected))
			}
			for i, expected := range test.expected {
				if !slicex.Equals(result[i], expected, IntEqual) {
					t.Fatalf("索引 %d: 实际: %+v, 预期: %+v", i, result[i], expected)
				}
			}
		})
	}
}

func TestSplitWithStrings(t *testing.T) {
	// 测试字符串类型的Split
	values := []string{"hello", "world", "test", "example", "demo"}
	var tests = []struct {
		name     string
		size     int
		index    int
		expected []string
	}{
		{
			name:     "字符串切片，size=2, index=0",
			size:     2,
			index:    0,
			expected: []string{"hello", "world"},
		},
		{
			name:     "字符串切片，size=2, index=2（最后一个不完整块）",
			size:     2,
			index:    2,
			expected: []string{"demo"},
		},
		{
			name:     "字符串切片，size=3, index=1（最后一个不完整块）",
			size:     3,
			index:    1,
			expected: []string{"example", "demo"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := slicex.Split(values, test.size)
			if test.index >= len(result) {
				t.Fatalf("索引 %d 超出结果切片长度 %d", test.index, len(result))
			}
			var actual = result[test.index]
			if !slicex.Equals(actual, test.expected, StringEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}

func TestSplitWithZeroSize(t *testing.T) {
	// 测试size=0的情况
	values := []int{1, 2, 3, 4, 5}
	result := slicex.Split(values, 0)
	if len(result) != 0 {
		t.Fatalf("size=0时应该返回空切片，实际: %+v", result)
	}
}
