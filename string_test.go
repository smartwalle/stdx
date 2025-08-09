package stdx_test

import (
	"testing"

	"github.com/smartwalle/stdx"
)

func TestString_Sub(t *testing.T) {
	tests := []struct {
		name     string
		s        stdx.String
		start    int
		length   int
		expected string
	}{
		{"正常情况", "Hello世界", 0, 5, "Hello"},
		{"包含中文", "Hello世界", 5, 2, "世界"},
		{"边界情况", "Hello世界", 0, 10, "Hello世界"},
		{"空字符串", "", 0, 5, ""},
		{"负起始位置", "Hello", -1, 3, ""},
		{"负长度", "Hello", 0, -1, ""},
		{"零长度", "Hello", 0, 0, ""},
		{"起始位置超出范围", "Hello", 10, 3, ""},
		{"长度超出范围", "Hello", 0, 10, "Hello"},
		{"起始位置等于字符串长度", "Hello", 5, 3, ""},
		{"emoji测试", "Hello👋世界", 5, 2, "👋世"},
		{"混合字符", "Hello👋世界", 0, 7, "Hello👋世"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.s.Sub(tt.start, tt.length)
			if result != tt.expected {
				t.Errorf("Sub(%q, %d, %d) = %q, 期望 %q", tt.s, tt.start, tt.length, result, tt.expected)
			}
		})
	}
}

func TestString_Between(t *testing.T) {
	tests := []struct {
		name     string
		s        stdx.String
		start    int
		end      int
		expected string
	}{
		{"正常情况", "Hello世界", 0, 4, "Hello"},
		{"包含中文", "Hello世界", 5, 6, "世界"},
		{"边界情况", "Hello世界", 0, 6, "Hello世界"},
		{"空字符串", "", 0, 5, ""},
		{"负起始位置", "Hello", -1, 3, ""},
		{"结束位置小于起始位置", "Hello", 3, 2, ""},
		{"起始位置等于结束位置", "Hello", 2, 2, "l"},
		{"起始位置超出范围", "Hello", 10, 12, ""},
		{"结束位置超出范围", "Hello", 0, 4, "Hello"},
		{"emoji测试", "Hello👋世界", 5, 6, "👋世"},
		{"混合字符", "Hello👋世界", 0, 7, "Hello👋世界"},
		{"单字符提取", "Hello", 1, 1, "e"},
		{"边界字符", "Hello", 0, 0, "H"},
		{"边界字符末尾", "Hello", 4, 4, "o"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.s.Between(tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("Between(%q, %d, %d) = %q, 期望 %q", tt.s, tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestString_Count(t *testing.T) {
	tests := []struct {
		name     string
		s        stdx.String
		expected int
	}{
		{"英文字符", "Hello", 5},
		{"中文字符", "世界", 2},
		{"混合字符", "Hello世界", 7},
		{"空字符串", "", 0},
		{"emoji", "👋🌍", 2},
		{"混合emoji", "Hello👋世界", 8},
		{"特殊字符", "Hello, 世界!", 10},
		{"数字", "12345", 5},
		{"空格", "   ", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.s.Count()
			if result != tt.expected {
				t.Errorf("Count(%q) = %d, 期望 %d", tt.s, result, tt.expected)
			}
		})
	}
}
