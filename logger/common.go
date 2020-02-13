package logger

import (
	"fmt"
	"runtime"
	"unsafe"
)

//日志级别
type Loglevel uint8

const (
	DEBUG Loglevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//获取当前运行函数的信息  skip表示展示的函数是第几层 skip=0表示 getInfo()函数 skip=1表示调用getInfo()函数的函数...
func getInfo(skip int) (file string, funcname string, line int) {
	//file文件全路径 line行号 ok调用是否出错
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	//file只要最后一段路径
	//file = path.Base(file)
	funcname = runtime.FuncForPC(pc).Name()
	return
}

/*
	字符串和字节切片相互转换
*/
func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
