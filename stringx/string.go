package stringx

import (
	"strings"
	"unicode/utf8"
)

// Sub 从指定位置开始提取指定长度的子字符串
func Sub(s string, start, length int) string {
	if s == "" || start < 0 || length <= 0 {
		return ""
	}
	var runes = []rune(s)
	if start >= len(runes) {
		return ""
	}
	var end = start + length
	if end > len(runes) {
		end = len(runes)
	}
	return string(runes[start:end])
}

// Between 提取两个位置之间的子字符串(闭合区间)
func Between(s string, start, end int) string {
	if s == "" || start < 0 || end < start {
		return ""
	}
	var runes = []rune(s)
	if start >= len(runes) {
		return ""
	}
	if end >= len(runes) {
		end = len(runes) - 1
	}
	return string(runes[start : end+1])
}

// Count 返回字符数量
func Count(s string) int {
	return utf8.RuneCountInString(s)
}

// Index 返回子字符串在字符串中第一次出现的位置
// 如果字符串为空或子字符串不存在，返回 -1
func Index(s string, substr string) int {
	if substr == "" {
		return -1
	}
	if s == "" {
		return -1
	}

	var index = strings.Index(s, substr)
	if index < 0 {
		return index
	}
	return utf8.RuneCountInString(s[:index])
}
