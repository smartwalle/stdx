package stdx_test

import (
	"github.com/smartwalle/stdx"
	"testing"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		v1 int
		v2 int
		r  int
	}{
		{0, 1, 1},
	}

	for _, tt := range tests {
		if actual := stdx.Max(tt.v1, tt.v2); actual != tt.r {
			t.Errorf("Max(%d, %d), 期望得到:%d, 实际得到:%d", tt.v1, tt.v2, tt.r, actual)
		}
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		v1 int
		v2 int
		r  int
	}{
		{0, 1, 0},
	}

	for _, tt := range tests {
		if actual := stdx.Min(tt.v1, tt.v2); actual != tt.r {
			t.Errorf("Min(%d, %d), 期望得到:%d, 实际得到:%d", tt.v1, tt.v2, tt.r, actual)
		}
	}
}
