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
