package stdx_test

import (
	"github.com/smartwalle/stdx"
	"math"
	"testing"
)

func TestToBool(t *testing.T) {
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
		if actual := stdx.ToBool(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 bool, 期望获得 %v, 实际获得  %v", tt.v, tt.r, actual)
		}
	}
}

func TestToFloat64(t *testing.T) {
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
		if actual := stdx.ToFloat64(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 float64, 期望获得 %f, 实际获得  %f", tt.v, tt.r, actual)
		}
	}
}

func TestToInt(t *testing.T) {
	var tests = []struct {
		v interface{}
		r int
	}{
		{1, 1},
		{9, 9},
		{9999, 9999},
		{uint64(9999), 9999},
		{int64(9999), 9999},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{-2, -2},
		{"1", 1},
		{"-1", -1},
		{"999999", 999999},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
	}

	for _, tt := range tests {
		if actual := stdx.ToInt(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 int, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestToInt64(t *testing.T) {
	var tests = []struct {
		v interface{}
		r int64
	}{
		{1, 1},
		{9, 9},
		{9999, 9999},
		{uint64(9999), 9999},
		{int64(9999), 9999},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{"1", 1},
		{"999999", 999999},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
		{"", 0},
		{"3414416614257328130", 3414416614257328130},
	}

	for _, tt := range tests {
		if actual := stdx.ToInt64(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 int64, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestToUint(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint
	}{
		{"1", 1},
		{"999999", 999999},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
	}

	for _, tt := range tests {
		if actual := stdx.ToUint(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint32, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestToUint8(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint8
	}{
		{"1", 1},
		{"65", 65},
		{"9.999", 9},
		{"99.1", 99},
		{"1.1111", 1},
		{"1.9999", 1},
		{100, 100},
		{99, 99},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(99), 99},
		{uint8(math.MaxUint8), math.MaxUint8},
	}

	for _, tt := range tests {
		if actual := stdx.ToUint8(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint8, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestToUint16(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint16
	}{
		{"1", 1},
		{"65535", 65535},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
		{uint16(math.MaxUint16), math.MaxUint16},
	}

	for _, tt := range tests {
		if actual := stdx.ToUint16(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint16, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestToUint32(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint32
	}{
		{"1", 1},
		{"999999", 999999},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
		{uint32(math.MaxUint32), math.MaxUint32},
	}

	for _, tt := range tests {
		if actual := stdx.ToUint32(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint32, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestToUint64(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint64
	}{
		{"1", 1},
		{"999999", 999999},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
		{uint64(math.MaxUint64), math.MaxUint64},
	}

	for _, tt := range tests {
		if actual := stdx.ToUint64(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 Uint64, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestToString(t *testing.T) {
	var tests = []struct {
		v interface{}
		r string
	}{
		{"1", "1"},
		{1, "1"},
		{1.1, "1.1"},
		{true, "true"},
		{3414416614257328130, "3414416614257328130"},
		{[]byte("Hello世界"), "Hello世界"},
		{[]rune("Hello世界"), "Hello世界"},
	}

	for _, tt := range tests {
		if actual := stdx.ToString(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 string, 期望获得 %v, 实际获得  %v", tt.v, tt.r, actual)
		}
	}
}
