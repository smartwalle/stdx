package stdx

import "unsafe"

// UnsafeBytes 将字符串转换成 byte 切片，注意不能对返回的切片进行修改
func UnsafeBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// UnsafeString 将 byte 切片转换成字符串，注意更新原切片的元素将会影响字符串
func UnsafeString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
