package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

/*
	性能测试
------------------------------
	go-torch 和火焰图
	go get -v github.com/uber/go-torch
	安装FlameGraph github上下载并将FlameGraph目录并添加到环境变量中
		先安装perl语言环境
		在修改go-torch/renderer/flamegraph.go
		*74行之后添加 这个修改后改回来才对,不确定操作的正确性
		if runtime.GOOS = "windows"{
			return  runScript("perl",append([]string{flameGraph},args...),graphInput)
		}
	在go install 就行了
------------------------------
压测工具wrk
github/wg/wrk 或者 github/adjust/go-wrk
------------------------------
	os.Args获取命令行参数 实际上是一个[]string 第0个参数是应用程序名
	但是不能解决 -name=xxx 这种参数的问题 就需要flag包
	pprof.exe -cpu=true 启动cpu测试
	go tool pprof cpu.pprof 查看cpu信息,进入一个交互式模式
	top 3 查看前三消耗cpu的函数
	web 用web图形化显示 需要安装 Graphviz 软件
--------------
Type: cpu
Time: Feb 20, 2020 at 2:49pm (CST)
Duration: 10.14s, Total samples = 36.53s (360.25%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top 3
Showing nodes accounting for 36.50s, 99.92% of 36.53s total
Dropped 11 nodes (cum <= 0.18s)
当前函数占用cpu耗时	百分比 总百分比 当前函数+调用当前函数的函数占用cpu耗时 百分比
      flat  flat%   sum%        cum   cum%
    17.94s 49.11% 49.11%     27.78s 76.05%  runtime.selectnbrecv
     9.84s 26.94% 76.05%      9.84s 26.94%  runtime.chanrecv
     8.72s 23.87% 99.92%     36.50s 99.92%  main.loginCode

*/
func loginCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("%v", v)
		default:
			time.Sleep(time.Millisecond * 5)
		}
	}
}
func main() {
	var isCPUPprof bool
	var isMemPprof bool
	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()
	//cpu性能消耗
	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Println("open file failed,err:", err)
			return
		}
		err = pprof.StartCPUProfile(file)
		if err != nil {
			fmt.Println("start cpu profile failed,err:", err)
			return
		}
		defer func() {
			pprof.StopCPUProfile()
			file.Close()
		}()
	}
	//测试函数
	for i := 0; i < 4; i++ {
		go loginCode()
	}
	time.Sleep(time.Second * 10)

	//内存性能消耗
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Println("open file failed,err:", err)
			return
		}
		err = pprof.WriteHeapProfile(file)
		if err != nil {
			fmt.Println("write mem profile failed,err:", err)
			return
		}
		defer func() {
			file.Close()
		}()
	}
}
