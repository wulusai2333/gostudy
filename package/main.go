package main

/*
	package
	使用包内函数 需要先导入包
	同一个文件夹下的文件都必须统一 package
		比如 package 目录下 main.go wlc.go
		main.go 设置package main | wlc.go 设置package wlc 就会报错
	import包需要从 GOPATH下src路径开始
	导入的包可以起别名
	go语言目前导入包的局限性:
		没有版本控制,由于go源码都必须放在GOPATH或者GOROOT的src目录下,导致如果包作者更新了包可能导致不兼容问题,而go包没有版本选择
	禁止循环导入包
		如: package main  import "wlc" | package wlc import "main"
	匿名导入包:
		import _ "数据库包"
		不使用包内部函数,但是包内init()函数会自动执行没有参数也没有返回值,

*/
import (
	"fmt"
	w "github.com/wulusai2333/gostudy/package/wlc"
)

func main() {
	fmt.Println(w.Add(1, 2))

}
