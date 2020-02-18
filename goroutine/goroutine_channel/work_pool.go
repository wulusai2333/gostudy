package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	计算随机数各位数的和
*/
type job struct {
	num int64
}
type result struct {
	*job
	sum int64
}

func sumNumber() {
	var wg sync.WaitGroup         //管理线程组的
	defer wg.Wait()               //等待组内线程,执行完才能关闭主线程
	var ch1 = make(chan *job, 10) //chan也能存结构体
	var ch2 = make(chan *result, 10)
	for i := 0; i < 24; i++ { //开启24个线程完成这个任务
		wg.Add(1)
		go add(&wg, ch1, ch2)
	}
	//死循环
	for {
		x := rand.Int63()
		ch1 <- &job{num: x} //存入结构体到通道1
		r := <-ch2          //从通道2取计算后的结果
		fmt.Println("num:", r.num, " sum:", r.sum)
		time.Sleep(time.Second)
	}
}

func add(wg *sync.WaitGroup, jChan <-chan *job, r chan<- *result) {
	defer wg.Done() //告诉组长咱执行完了
	for {
		j := <-jChan //从通道1取job结构体
		n := j.num
		var sum int64
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		r <- &result{ //计算完的结构体存进通道2
			job: j,
			sum: sum,
		}
	}
}
