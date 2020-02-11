package wlc

import "fmt"

/*
	与之前的规则一样只有首字母大写才能被外界调用
*/
func Add(a int, b int) int {
	return a + b
}

var x int = 100
var y = 10

/*
	init()的执行时机 全局变量先生成 再执行init() 最后才执行main()
	如果 package a import b
	package b import c
	则执行顺序 c.init() -> b.init() -> a.init() -> main.init()
*/
func init() {
	fmt.Println("import 时自动执行")
	fmt.Println(x, y)
}
