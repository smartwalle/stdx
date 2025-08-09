package stdx

import (
	"unicode/utf8"
)

type String string

// Sub 从指定位置开始提取指定长度的子字符串
func (s String) Sub(start, length int) string {
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
func (s String) Between(start, end int) string {
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
func (s String) Count() int {
	return utf8.RuneCountInString(string(s))
}
