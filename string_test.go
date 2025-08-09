package stdx_test

import (
	"github.com/smartwalle/stdx"
	"testing"
)

func TestString(t *testing.T) {
	var tests = []struct {
		v interface{}
		r string
	}{
		{"1", "1"},
		{1, "1"},
		{1.1, "1.1"},
		{true, "true"},
		{3414416614257328130, "3414416614257328130"},
	}

	for _, tt := range tests {
		if actual := stdx.String(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 string, 期望获得 %v, 实际获得  %v", tt.v, tt.r, actual)
		}
	}
}
