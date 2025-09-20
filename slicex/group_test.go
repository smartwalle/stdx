package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"testing"
)

func TestGroup(t *testing.T) {
	var values = []int{35, 34, 33, 32, 31, 1, 2, 3, 4, 5, 11, 12, 13, 14, 15, 21, 22, 23, 24, 25}
	var actual = slicex.Group(values, func(elem int) int {
		return elem % 10
	})
	t.Log(actual)
}

func TestGroupMatched(t *testing.T) {
	var values = []int{35, 34, 33, 32, 31, 1, 2, 3, 4, 5, 11, 12, 13, 14, 15, 21, 22, 23, 24, 25}
	var actual = slicex.GroupMatched(values, func(elem int) bool {
		return elem%10 != 1
	}, func(elem int) int {
		return elem % 10
	})
	t.Log(actual)
}
