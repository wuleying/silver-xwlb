
package utils

import "unsafe"

// StringSub 截断字符串
func StringSub(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// Str2bytes 字符串转byte
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}

	return *(*[]byte)(unsafe.Pointer((&h)))
}

// Bytes2str byte转字符串
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}