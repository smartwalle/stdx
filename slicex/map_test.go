package slicex_test

import (
	"github.com/smartwalle/stdx/slicex"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	var tests = []struct {
		source []int
		expect []string
		fn     func(elem int) string
	}{
		{source: []int{1, 2, 3, 4, 5, 6}, expect: []string{"2", "3", "4", "5", "6", "7"}, fn: func(elem int) string {
			return strconv.FormatInt(int64(elem+1), 10)
		}},
		{source: []int{1, 2, 3, 4, 5, 6}, expect: []string{"2", "4", "6", "8", "10", "12"}, fn: func(elem int) string {
			return strconv.FormatInt(int64(elem*2), 10)
		}},
	}

	for _, test := range tests {
		var actual = slicex.Map(test.source, test.fn)
		if !slicex.Equals(actual, test.expect, StringEqual) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}

func TestMapMatched(t *testing.T) {
	var tests = []struct {
		source     []int
		expect     []string
		filterFunc func(elem int) bool
		mapFunc    func(elem int) string
	}{
		{source: []int{1, 2, 3, 4, 5, 6}, expect: []string{"1", "2", "3", "4", "5", "6"}, filterFunc: func(elem int) bool {
			return true
		}, mapFunc: func(elem int) string {
			return strconv.FormatInt(int64(elem), 10)
		}},
		{source: []int{1, 2, 3, 4, 5, 6}, expect: []string{"2", "4", "6", "8", "10", "12"}, filterFunc: func(elem int) bool {
			return true
		}, mapFunc: func(elem int) string {
			return strconv.FormatInt(int64(elem*2), 10)
		}},
		{source: []int{1, 2, 3, 4, 5, 6}, expect: []string{"2", "4", "6"}, filterFunc: func(elem int) bool {
			return elem%2 == 0
		}, mapFunc: func(elem int) string {
			return strconv.FormatInt(int64(elem), 10)
		}},
	}

	for _, test := range tests {
		var actual = slicex.MapMatched(test.source, test.filterFunc, test.mapFunc)
		if !slicex.Equals(actual, test.expect, StringEqual) {
			t.Fatalf("实际: %+v, 预期: %+v", actual, test.expect)
		}
	}
}
