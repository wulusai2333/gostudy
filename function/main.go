package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	函数
	func 函数名(参数)(返回值){
		函数体
	}
	函数分为如下几种:
		无参无返回值/无参有返回值/有参有返回值/有参无返回值
	说白了就是参数和返回值都可以省略
	而且与java语言不同,go可以多返回值,返回值可以用 _ 丢弃不用
	多返回值,命名返回值需要用()括起来
*/
/*
	内置函数
	close 主要用来关闭channel
	len 求长度 string array slice map channel
 	new 分配内存 主要分配值类型 如 int struct 返回的指针
 	make 分配内存 主要分配引用类型 chan map slice
	append 追加数组元素到数组,slice中
	panic和recover 用来错误处理
*/
func main() {
	var arr [10]string
	fmt.Println(arr)
	sum4()
	hanzi("sdjj是的我发给")
	danci("how do you do")
	huiwen("上海自来水来自海上")
	_defer()
	//defer 的返回时机
	fmt.Println(der1())
	fmt.Println(der2())
	fmt.Println(der3())
	fmt.Println(der4())
	//变量的作用域
	f2()
	//函数作为参数
	fun2(fun1)
	fun3()
	fun4(fun1)
	//闭包的使用
	bb4()
	addr3()
	test()
}

/*
	返回值命名和不命名区别,命名的返回值可以直接在函数中使用,相比return更灵活
	相当于提前声明了变量
*/

func sum(a int, b int) (ret int) {

	ret = a + b
	return
}

//匿名返回值不需要括号
func sum1(int, int) int {
	return 1
}

//多返回值需要括号 多返回值要么都命名要么都不命名
func sum2(int, int) (int, int) {
	return 1, 1
}

//参数的简写
func sum3(x, y int, a int) int {
	return x + y
}

//可变参数 填入的参数数量可以为0个或多个
func sum4(...int) int {
	return 0
}

/*
	判断字符串中汉字的数量
*/
func hanzi(str string) (num int) {
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			num++
		}
	}
	fmt.Println("汉字数量:", num)
	return
}

/*
	判断单词出现的次数,并统计存入map
*/
func danci(str string) {
	m := make(map[string]int)
	s := strings.Split(str, " ")
	for _, v := range s {
		m[v] += 1
	}
	fmt.Println(m)
}

/*
	回文判断
*/
func huiwen(str string) {
	s := strings.Split(str, "")
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println(s, ":是回文")
	return
}

/*
	defer 后面跟着的语句延迟处理 类似于压栈的操作,先进后出
	执行的时机
	在函数中使用defer时,有返回值时 , return x会经历如下几步,先给返回值x赋值 再执行defer的语句 最后才执行ret指令
*/
func _defer() {
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

//以下面试题 考察defer对返回值的影响 实际写代码没人这么写
func der1() int {
	x := 5
	defer func() {
		x++
	}()
	return x //5
}

func der2() (x int) {
	x = 5
	defer func() {
		x++
	}()
	return 5 //6
}

func der3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //5
}

func der4() (x int) {

	defer func(x int) {
		x++
	}(x) //这里就体现了go语言函数的值拷贝,里面的x相当于局部变量,不影响外面的x
	return 5 //5
}

/*
	变量的作用域 就近原则,跟java一样
*/

var p = 10

func f1(p int) {
	p = 12
	fmt.Println(p)
}
func f2() {
	p = 13901215948
	f1(p)
}

/*
	函数作为参数和返回值
	其实,函数可以做参数这点,本身规则没有变
*/
func fun1() int {
	return 1
}

func fun2(f func() int) {
	a := fun1
	fmt.Printf("%T\n", a)
}
func fun3() {
	a := fun2
	fmt.Printf("%T\n", a)
}
func fun4(f func() int) (y func() int) {
	fmt.Printf("%T\n", fun4)
	y = fun1
	return y
}

/*
	匿名函数
	可以用全局变量接收,但是一般不这么用, 通常用在函数内部
*/
var ff = func() {

}

func f3() {
	ff()
}

func f4() {
	//定义一个匿名函数并使用,匿名函数一般在函数内部使用的函数
	var f = func() {

	}
	f()
	//立即执行函数,某位多了对括号 对只用一次的函数这么操作
	func() {
		fmt.Printf("hello")
	}()
}

/*	闭包 一个函数及相关引用组合的实体   闭包=函数+引用环境
	闭包:对于参数不匹配函数的调用,通常使用闭包来解决
	场景: bb1是你同事写的代码,而你无权修改,但你又需要将bb2传到bb1里

	这个例子里 bb1 就是需要接收bb2作为参数的函数, bb3就是封装bb2的函数

*/

func bb1(f func()) {
	fmt.Println("is bb1")
	f()
}
func bb2(x, y int) {
	fmt.Println("is bb2")
	fmt.Println(x + y)
}
func bb3(f func(int, int), x, y int) func() {
	fmt.Println("is bb3")
	ff := func() {
		f(x, y)
	}
	return ff
}
func bb4() {
	fmt.Println("闭包例子:")
	f := bb3(bb2, 1, 2)
	fmt.Println("这里bb2还没运行")
	bb1(f)
	//函数运行的先后顺序 bb3 -> bb1 -> bb2 这就是闭包的意义所在
}

/*
	下面演示了闭包
	addr()返回值是一个方法
	addr3中 变量 f 接收的是方法addr()的返回值也就是一个函数
	f = func(i int) int {
		return x + i
	}
	f1接收的是 f (其实本身是一个函数) 返回的参数 return x + i

	闭包是什么? 一个函数,包含了这个函数外部的变量 这个例子中 f 就是一个闭包
	闭包利用的底层原理:
	1.函数可以作为返回值
	2.函数查找参数的顺序,现在自己内部找,找不到再往上层找
*/

func addr() func(int) int {
	var x int
	x = 10
	return func(i int) int {
		return x + i
	}
}

/*func addr2(x int) func(int) int {

}*/
func addr3() {
	//闭包的使用
	f := addr()
	f1 := f(100)
	fmt.Println(f1)
}

func test() {
	//txtFunc就是一个引入了外部参数的函数,感觉有点像java的工厂模式
	txtFunc := addSuffix(".txt")
	fmt.Println(txtFunc("xiaomi"))
	fmt.Println(txtFunc("xiaomi.txt"))
	//
	f1, f2 := calc(100)
	fmt.Println("calc执行结果:", f1(15), f2(50)) //115 65
}

/*
	给文件名加后缀
*/
func addSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

/*
	校验同步 返回的函数add 和 sub 公用了参数base 最后返回的结果函数 f1 f2 给他们分别传参 最终得到一个同步的效果,运行一个函数会影响另一个函数的base值
*/
func calc(base int) (func(int) int, func(int) int) {
	add := func(x int) int {
		base += x
		return base
	}
	sub := func(x int) int {
		base -= x
		return base
	}
	return add, sub
}
