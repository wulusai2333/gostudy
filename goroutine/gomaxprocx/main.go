package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	GOMAXPROCX
	go 语言的goroutine是轻量级的 一个初始栈的大小约2kb 最大支持1G
	而调用操作系统线程固定为2M 所以启动go线程消耗比调用操作系统线程开销小的多
调度模型 GMP
G: goroutine
M: 映射操作系统线程
P: 调度者,把G放在M上运行  M:N 把M个goroutine分配给N个操作系统线程执行
*/
func a() {
	defer wg.Done()
	for i := 1; i < 1000000; i++ {
		fmt.Println("A:", i)
	}
}
func b() {
	defer wg.Done()
	for i := 1; i < 1000000; i++ {
		fmt.Println("B:", i)
	}
}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)                //指定使用的cpu核心数 通常是小于等于cpu核心数的 这个程序单核跑,以免影响其他程序运行
	runtime.GOMAXPROCS(runtime.NumCPU()) //指定使用的cpu核心数为cpu总核心数,也就是跑满cpu 默认即这个值
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
