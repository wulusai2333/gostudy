package main

import (
	"fmt"
	"image"
	"sync"
	"time"
)

/*
	锁的使用
	互斥锁	只有一个线程能够访问资源
	读写分离的 读写互斥锁 读的时候都不涉及修改资源,所以所有线程都可以读,写的时候就将资源锁定
	在高读并发场景下,读写分离锁的效率远高于互斥锁
	sync.Once 多线程并发场景下只执行一次

*/
var (
	n      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex   //互斥锁,只有一个线程能访问资源
	rwLock sync.RWMutex //读写分离锁
)

func add() {
	for i := 1; i <= 5000000; i++ {
		lock.Lock()
		n += 1
		lock.Unlock()
	}
	wg.Done()
}
func write() {
	//lock.Lock()
	rwLock.Lock()
	n += 1
	time.Sleep(time.Millisecond * 5)
	rwLock.Unlock()
	//lock.Unlock()
	wg.Done()
}
func read() {
	//lock.Lock()
	rwLock.RLock()
	//fmt.Println(n)
	if n == 0 {
	}
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	//lock.Unlock()
	wg.Done()
}
func main() {
	//互斥锁测试
	lockTest()
	//读写分离锁测试
	readAndWriteLock()
	//内置map不安全
	unsafeMap()
	//sync包并发安全Map
	safeMap()
}

//内置map不支持并发
func unsafeMap() {
	var m = make(map[int]int)
	get := func(key int) int {
		return m[key]
	}
	set := func(key int, value int) {
		m[key] = value
	}
	for i := 0; i < 20; i++ { //不加锁内置map会报错 fatal error: concurrent map writes
		wg.Add(1)
		go func(n int) {
			key := n
			lock.Lock()
			set(key, n)
			lock.Unlock()
			fmt.Println(key, ":", get(key))
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

//sync包map
func safeMap() {
	var sm = sync.Map{}
	for i := 0; i < 20; i++ { //不加锁内置map会报错 fatal error: concurrent map writes
		wg.Add(1)
		go func(n int) {
			key := n
			sm.Store(key, n)
			value, _ := sm.Load(key)
			fmt.Println(key, ":", value)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

func lockTest() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(n)
}

func readAndWriteLock() {
	start := time.Now()
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go write()
	}
	for i := 1; i <= 5000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
	//253.3224ms 互斥锁9.1007977s 读写分离锁318.187ms
}

/*
	场景:服务启动时有些资源加载慢且不是启动必要的,可以先不加载,等需要访问时再加载
	在并发访问的情况下,可能造成某个资源多次加载或者加载到一半就被另外的进程访问了
sync.Once 接收参数只能是一个无参无返回值的参数,这时候就可以用闭包将有参函数包裹起来传递进去
*/
var icons map[string]image.Image
var once sync.Once

func loadIcons() {
	icons = map[string]image.Image{}
}

func icon(name string) image.Image {
	once.Do(loadIcons) //此资源只需加载一次
	return icons[name]
}
