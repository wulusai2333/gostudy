package utils

import "unsafe"

/*
	字符串和字节切片相互转换
*/
//字符串转字节数组
func Str2bytes(s string) []byte {
	return str2bytes(s)
}

//字节数组转字符串
func Bytes2str(b []byte) string {
	return bytes2str(b)
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

/*
	由于是非复制转类型,使用不当会发生不可期的错误,特别是数据会被map slice 等引用类型使用的时候
golang中分为值类型和引用类型
值类型分别有：int系列、float系列、bool、string、数组和结构体
引用类型有：指针、slice切片、管道channel、接口interface、map、函数等
*/
func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
