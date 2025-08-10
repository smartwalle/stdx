package stdx

import "unsafe"

// UnsafeBytes 将字符串转换成 byte 切片，注意不能对返回的切片进行修改
func UnsafeBytes(s string) []byte {
	if s == "" {
		return []byte{}
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// UnsafeString 将 byte 切片转换成字符串，注意更新原切片的元素将会影响字符串
func UnsafeString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(&b[0], len(b))
}
