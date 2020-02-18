package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	原子操作
*/
var (
	a    int
	b    int
	c    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

func addNoLock() {
	a++
	wg.Done()
}
func addWithLock() {
	lock.Lock()
	b++
	lock.Unlock()
	wg.Done()
}
func atom() {
	atomic.AddInt64(&c, 1)
	wg.Done()
}
func main() {
	//无锁 2.9913ms
	noLock()
	//互斥锁 2.9918ms
	haveLock()
	//原子相加
	addAtom()
	//改变值
	compareAdd()
}

func compareAdd() {
	var n int64
	ok := atomic.CompareAndSwapInt64(&n, 0, 200) //将一个变量的指针传过来跟old比较,如果两个值相等则将n位置的值替换为new
	fmt.Println(n, ok)
}
func addAtom() {
	start := time.Now()
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go atom()
	}
	wg.Wait()
	fmt.Println(c, time.Now().Sub(start))
}

func haveLock() {
	startLock := time.Now()
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go addWithLock()
	}
	wg.Wait()
	fmt.Println(b, time.Now().Sub(startLock))
}

func noLock() {
	start := time.Now()
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go addNoLock()
	}
	wg.Wait()
	fmt.Println(a, time.Now().Sub(start))
}
