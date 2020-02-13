package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"unsafe"
)

/*
	日志输出到文件
*/
func main() {
	file, err := os.OpenFile("log/logger/a.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	for {
		t := time.Now().Format("2006/01/02 15:04:05.000 ")
		str := "this is log for test\n"
		log.Print(t, str)
		_, err := file.Write(append(str2bytes(t), str2bytes(str)...))
		if err != nil {
			fmt.Println("write file failed, err:", err)
		}
		log.SetOutput(file) //这条就设置了log的打印位置 直接打印在文件中而不在终端输出
		time.Sleep(3 * time.Second)
	}

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
