package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

/*
	go 的线程启动方式
	goroutine启动只需要在函数前加 go 关键字
*/
var wg sync.WaitGroup

func main() {
	//使用strconv包对字符串转换操作
	_strconv()
	//循环启动多线程
	_goroutine()
	//rand包 使用时要指定一个随机数种子,以确保每次执行时随机的结果不一样 比如当前时间纳秒
	_rand()

	for i := 0; i < 10; i++ {
		wg.Add(i)   //线程计数+1
		go wgFun(i) //开启线程
	}
	wg.Wait() //等待线程执行完
}
func wgFun(i int) {
	defer wg.Done()                                        //表示线程已经执行完了
	time.Sleep(time.Second * time.Duration(rand.Intn(10))) //睡一会
	fmt.Println("wgFun:", i)
}
func _rand() {
	rand.Seed(1)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d %d\n", rand.Int(), rand.Intn(10))
	}
}

func _goroutine() {
	for i := 0; i < 10000; i++ {
		//go hello(i)
		//匿名函数版 如果不设参数则 i 用的是函数外部的参数,而外面的参数变了,go线程再去取值结果就可能出现重复
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}
func hello(i int) {
	fmt.Println(i)
}

func _strconv() {
	str1 := strconv.Itoa(100)
	//int 转 string
	fmt.Printf("%#v,%T\n", str1, str1)
	int2, _ := strconv.Atoi("10")
	//string 转 int
	fmt.Printf("%#v,%T\n", int2, int2)
	int6, _ := strconv.ParseInt("20", 10, 64)
	//string 转 int,int8,int64... 第二个参数是进制,第三个参数是比特位
	fmt.Printf("%#v,%T\n", int6, int6)
	float6, _ := strconv.ParseFloat("100.1001", 64)
	//string 转 float
	fmt.Printf("%#v,%T\n", float6, float6)
	boolean, _ := strconv.ParseBool("true")
	//string 转 bool
	fmt.Printf("%#v,%T\n", boolean, boolean)
}
