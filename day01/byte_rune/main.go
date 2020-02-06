package main

import "fmt"

/*
	go语言的流程控制

	if

	for

	switch

	goto
*/

func main() {
	_if()
	_for()
	_switch()
	_goto()
}

func _if() {
	//if语句可以在判断条件时定义变量
	if age := 18; age > 19 {
		fmt.Println("if...")
	} else {
		fmt.Println("else...")
	}
}

func _for() {
	//for循环的几种形式
	//基本格式
	for i := 1; i > 5; i++ {
		fmt.Println(i)
	}
	//省略初始语句
	var a = 1
	for ; a < 5; a++ {
		fmt.Println(a)
	}
	//省略步进
	for a < 10 {
		fmt.Println(a) //a=5
		break          //死循环跳出循环
	}
	//无限循环
	for {
		fmt.Println()
		break
	}
	/* for range 遍历键值 :数组 切片 字符串 map 通道
	数组 切片 字符串 返回索引和值
	map 返回键值
	通道 只返回值
	*/
	s := "hello我的" //汉字占三个字节
	for k, v := range s {
		fmt.Printf("%d %c\n", k, v)
		if k == 8 {
			continue //跳过当前次循环 与java语法相同
		}
	}
}

/*
	在实际应用上 switch 简化了大量if else if的判断
*/
func _switch() {
	//基本写法
	a := 5
	switch a {
	case 1:
		fmt.Println(a)
	case 2:
		fmt.Println(a)
	case 3:
		fmt.Println(a)
	case 4:
		fmt.Println(a)
	case 5:
		fmt.Println(a)
	default:
		fmt.Println("null")
	}
	//变量定义在函数中 同一个case包含多个值

	switch b := 6; b {
	case 1, 2, 3, 4:
		fmt.Println("小于5")
	case 5, 6, 7, 8, 9:
		fmt.Println("不小于5")
	default:
		fmt.Println("过不来")
	}

	//case后判断
	switch c := 4; {
	case c < 1:
		fmt.Println("c小于1")
		fallthrough //为了兼容c语言设计的 可以穿透到下一个case
	case c > 5:
		fmt.Println("c大于5")
	default:
		fmt.Println("不小也不大")
	}

}

/*
	goto 代码里定义了一些标志位 goto可以跳转到对应位置
	由于goto跳转的特性 在某些特定场景会很有用 但是如果代码里加了大量goto语句就会让代码很难懂

*/
func _goto() {

	for i := 0; i < 10; i++ {
		fmt.Println("goto ", i)
		if i == 5 {
			goto tag
		}
	}
	return
	//标签
tag:
	fmt.Println("标签")
}
