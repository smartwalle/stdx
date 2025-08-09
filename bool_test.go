package stdx_test

import (
	"github.com/smartwalle/stdx"
	"testing"
)

func TestBool(t *testing.T) {
	var tests = []struct {
		v interface{}
		r bool
	}{
		{"true", true},
		{"t", true},
		{"1", true},
		{"0", false},
		{"2", false},
		{"false", false},
		{0, false},
		{2, true},
		{-2, true},
		{1, true},
	}

	for _, tt := range tests {
		if actual := stdx.Bool(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 bool, 期望获得 %v, 实际获得  %v", tt.v, tt.r, actual)
		}
	}
}
