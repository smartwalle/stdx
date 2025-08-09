package stdx_test

import (
	"github.com/smartwalle/stdx"
	"testing"
)

func TestFloat64(t *testing.T) {
	var tests = []struct {
		v interface{}
		r float64
	}{
		{"9.9", 9.9},
		{"9.99", 9.99},
		{"9.91", 9.91},
		{"9.90", 9.90},
		{"9.901", 9.901},
		{"9.001", 9.001},
	}

	for _, tt := range tests {
		if actual := stdx.Float64(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 float64, 期望获得 %f, 实际获得  %f", tt.v, tt.r, actual)
		}
	}
}
