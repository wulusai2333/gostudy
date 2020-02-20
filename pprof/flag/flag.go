package main

import (
	"flag"
	"fmt"
	"time"
)

/*
	flag的用法
	-name="小明"
	-name 小明
	--name 小明 都可以
*/
func main() {
	//三个参数分别为 flag名 默认参数 提示
	name := flag.String("name", "小呜", "名字")
	flag.Parse()
	fmt.Println(*name, time.Now().Format("2006-01-02 03:04:05"))
	fmt.Println(flag.Args())  //命令行参数后的其他参数
	fmt.Println(flag.NArg())  //命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) //返回使用命令行参数的个数
}
