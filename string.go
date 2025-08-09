package stdx

import (
	"strings"
	"unicode/utf8"
	"unsafe"
)

// String 以字符为单位处理字符串
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

// Index 返回子字符串在字符串中第一次出现的位置
// 如果子字符串为空，返回 0
// 如果字符串为空或子字符串不存在，返回 -1
func (s String) Index(substr string) int {
	if substr == "" {
		return 0
	}
	if s == "" {
		return -1
	}

	var index = strings.Index(string(s), substr)
	if index < 0 {
		return index
	}
	return utf8.RuneCountInString(string(s[:index]))
}

// Bytes 转换成 byte 切片，注意不能对返回的切片进行修改
func (s String) Bytes() []byte {
	return unsafe.Slice(unsafe.StringData(string(s)), len(s))
}
