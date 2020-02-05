package main

import "fmt"

/*
	go语言执行从main函数开始 如果想单独执行必须得是main包下的main函数
	其他包将作为工具被引用,也不能直接执行

	go build的使用
	当前目录下执行go build 会编译当前目录下的代码为可执行文件 windows为.exe linux为二进制文件
	go build后还可以跟路径 路径为GOPATH下的路径 如go build wulusai/hello 就会去编译$GOPATH/src/wulusai/hello目录下的文件
	而go build编译时不会管下级目录的main函数
	后面加 -o 参数 可以定义编译后二进制文件的名字

	go run 直接运行程序(不推荐使用)
	go install 包含两步,先编译二进制文件再将二进制文件放入到$GOPATH/bin目录下
	go编译直接生成二进制文件的好处: 运行不需要安装go环境,只要对应平台的二进制文件即可运行

	go语言可以跨平台编译 即在windows平台下编译可以在linux平台运行的二进制文件

	SET CGO_ENABLE=0  //禁用CGO 因为跨平台编译会出现问题
	SET GOOS=linux  //设定目标操作平台
	SET GOARCH=amd64 //目标处理器架构amd64

	linux,windows,mac平台下编译跨平台设置各有不同,这里不再说明,需要的时候再去找

	go语言 package关键字 声明代码所在的包 如果为main包 最终会生成一个可执行的二进制文件
*/
func main() {
	fmt.Print("hello world")            //打印不换行
	fmt.Printf("world %s hello", "123") //带占位符
	fmt.Println("hello")                //换行打印
	//这部分总体语法跟java差不多
}
