package main

import "fmt"

/*
	struct 结构体
	type 类型
*/
type myint int //自定义类型	显示的是 main.myint 类型
type me = int  //类型别名,只在当前生效,实际还是 int 类型
func main() {
	fmt.Println(f(5))
	fmt.Println(taijie(10))
	var a myint = 100
	var b me = 100
	fmt.Printf("myint:%v,%T\n", a, a)
	fmt.Printf("me:%v,%T\n", b, b)
}

/*
	阶层
*/
func f(n uint) uint {
	if n == 1 {
		return 1
	}
	return n * f(n-1)
}

/*
	上台阶 n阶台阶有多少种走法
*/
func taijie(n uint) uint {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}
