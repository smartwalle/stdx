package slicex_test

import (
	"testing"

	"github.com/smartwalle/stdx/slicex"
)

func TestEach(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		expected []int
	}{
		{
			name:     "正常切片",
			source:   []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "空切片",
			source:   []int{},
			expected: []int{},
		},
		{
			name:     "单元素切片",
			source:   []int{42},
			expected: []int{42},
		},
		{
			name:     "nil切片",
			source:   nil,
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result []int
			slicex.Each(test.source, func(elem int) {
				result = append(result, elem)
			})

			if !slicex.Equals(result, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", result, test.expected)
			}
		})
	}
}

func TestReverseEach(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		expected []int
	}{
		{
			name:     "正常切片",
			source:   []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "空切片",
			source:   []int{},
			expected: []int{},
		},
		{
			name:     "单元素切片",
			source:   []int{42},
			expected: []int{42},
		},
		{
			name:     "nil切片",
			source:   nil,
			expected: []int{},
		},
		{
			name:     "两个元素切片",
			source:   []int{1, 2},
			expected: []int{2, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result []int
			slicex.ReverseEach(test.source, func(elem int) {
				result = append(result, elem)
			})

			if !slicex.Equals(result, test.expected, IntEqual) {
				t.Fatalf("实际: %+v, 预期: %+v", result, test.expected)
			}
		})
	}
}

func TestEachWithModification(t *testing.T) {
	// 测试在Each过程中修改元素的行为
	source := []int{1, 2, 3, 4, 5}
	var sum int
	slicex.Each(source, func(elem int) {
		sum += elem
	})

	expected := 15 // 1+2+3+4+5
	if sum != expected {
		t.Fatalf("实际: %+v, 预期: %+v", sum, expected)
	}
}

func TestReverseEachWithModification(t *testing.T) {
	// 测试在ReverseEach过程中修改元素的行为
	source := []int{1, 2, 3, 4, 5}
	var sum int
	slicex.ReverseEach(source, func(elem int) {
		sum += elem
	})

	expected := 15 // 5+4+3+2+1
	if sum != expected {
		t.Fatalf("实际: %+v, 预期: %+v", sum, expected)
	}
}
