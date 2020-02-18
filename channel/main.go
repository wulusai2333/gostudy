package main

import (
	"fmt"
	"github.com/wulusai2333/gostudy/channel/logger"
	"sync"
)

/*
	go语言并发模型是CSP 通过通信共享内存  而不是 通过共享内存通信
	channel 遵循先进先出
*/
//声明通道 这个表示将一个int类型的数据放进通道内
var c chan int
var wg sync.WaitGroup
var once sync.Once //确保某个操作只执行一次
func main() {
	//noBufChan() //无缓冲
	//bufChan()   //有缓冲
	//useChan()
	log := logger.NewFileLogger(logger.INFO, "channel/logger/", "chanFile", 1024*1024)
	for {
		log.Debug("this is 异步打印日志")
		log.Info("this is 异步打印日志")
		log.Error("this is 异步打印日志")
		//time.Sleep(time.Second / 500)
	}
}

/*
	单项通道 ch1只能取值不能存值,ch2只能存值不能取值
	通常作为函数的参数使用
*/
func danxiang(ch1 <-chan int, ch2 chan<- int) {
	//ch1<-1 //不能存
	<-ch1
	//<-ch2 //不能取
	ch2 <- 1
}
func useChan() {
	defer wg.Wait()
	wg.Add(3)
	c1 := make(chan int, 20)
	c2 := make(chan int, 100)
	go f1(c1)
	go f2(c1, c2)
	go f2(c1, c2)
	for {
		n, ok := <-c2
		if !ok {
			break
		}
		fmt.Println("c2取出的值:", n)
	}
}

func f1(c1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		c1 <- i
	}
	close(c1)
}

func f2(c1, c2 chan int) {
	defer wg.Done()
	for {
		n, ok := <-c1
		if !ok {
			break
		}
		c2 <- n * n
		fmt.Println("c1里的值:", n) //这条语句如果方面存值前面会发生send to closed channel错误
	}
	once.Do(func() {
		close(c2) //close只执行一次
	})
}

//有缓冲通道 存满之前程序可以继续执行下去
func bufChan() {
	c = make(chan int, 10)
	//通道的初始化  带缓冲区的通道
	/*
		有无缓冲区的区别 无缓冲,放一个后面的都得等待这个被拿走其他人才能放
		有缓冲区,放到缓冲区上限之前都不用等
	*/
	//通道存值 通道里尽量存小值 如果值比较大就存地址,如string字符串就放地址
	c <- 10
	//通道取值
	x := <-c
	//<-c 也可以不接收 表示把这个数据扔了
	fmt.Println(x)
	//关闭
	close(c)
}

//无缓冲通道 存了不取会产生死锁
func noBufChan() {
	defer wg.Wait()
	wg.Add(1)
	c = make(chan int)
	//通道的初始化 大小可以不指定 无缓冲区的通道
	go func() {
		fmt.Println("值被其他线程取出来了", <-c) //打印放在通道的值
		defer wg.Done()
	}()
	//另一个线程把卡在通道的值取出来了
	c <- 1
	//无缓冲区,如果没有其他线程接收这个值就会产生死锁
	fmt.Println("把值存到通道里了")
	//<-c                    //想丢,然而在上一步就卡主了
	close(c)
}
