package slicex_test

import (
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestReduce(t *testing.T) {
	tests := []struct {
		name     string
		source   []int
		initial  int
		fn       func(int, int) int
		expected int
	}{
		{
			name:     "求和",
			source:   []int{1, 2, 3, 4},
			initial:  0,
			fn:       func(acc, x int) int { return acc + x },
			expected: 10,
		},
		{
			name:     "nil切片",
			source:   []int{},
			initial:  100,
			fn:       func(acc, x int) int { return acc + x },
			expected: 100,
		},
		{
			name:     "统计元素数量",
			source:   []int{1, 2, 3},
			initial:  0,
			fn:       func(acc int, _ int) int { return acc + 1 },
			expected: 3,
		},
		{
			name:    "求最小值",
			source:  []int{5, 2, 8, 1, 9},
			initial: 5,
			fn: func(acc, x int) int {
				if x < acc {
					return x
				}
				return acc
			},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := slicex.Reduce(test.source, test.initial, test.fn)
			if actual != test.expected {
				t.Errorf("实际: %v, 预期： %v", actual, test.expected)
			}
		})
	}
}

func TestReduceWithStrings(t *testing.T) {
	tests := []struct {
		name     string
		source   []string
		initial  string
		fn       func(string, string) string
		expected string
	}{
		{
			name:     "拼接字符串",
			source:   []string{"a", "b", "c"},
			initial:  "",
			fn:       func(acc, x string) string { return acc + x },
			expected: "abc",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := slicex.Reduce(test.source, test.initial, test.fn)
			if actual != test.expected {
				t.Errorf("实际: %v, 预期: %v", actual, test.expected)
			}
		})
	}
}

func TestReduceWithMultiStats(t *testing.T) {
	type stats struct {
		min   int
		max   int
		sum   int
		count int
	}

	source := []int{5, 2, 8, 1, 9}
	initial := stats{min: source[0], max: source[0], sum: 0, count: 0}
	actual := slicex.Reduce(source, initial, func(acc stats, x int) stats {
		if x < acc.min {
			acc.min = x
		}
		if x > acc.max {
			acc.max = x
		}
		acc.sum += x
		acc.count++
		return acc
	})

	expected := stats{min: 1, max: 9, sum: 25, count: 5}
	if actual != expected {
		t.Errorf("实际: %+v, 预期: %+v", actual, expected)
	}
}
