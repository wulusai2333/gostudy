package main

import (
	"fmt"
	"strings"
)

/*
	go语言执行从main函数开始 如果想单独执行必须得是main包下的main函数
	其他包将作为工具被引用,也不能直接执行

	go build的使用
	当前目录下执行go build 会编译当前目录下的代码为可执行文件 windows为.exe linux为二进制文件
	go build后还可以跟路径 路径为GOPATH下的路径 如go build wulusai/hello 就会去编译$GOPATH/src/wulusai/hello目录下的文件
	而go build编译时不会管下级目录的main函数
	后面加 -o 参数 可以定义编译后二进制文件的名字

	go run 直接运行程序(不推荐使用)
	go install 包含两步,先编译二进制文件再将二进制文件放入到$GOPATH/bin目录下
	go编译直接生成二进制文件的好处: 运行不需要安装go环境,只要对应平台的二进制文件即可运行

	go语言可以跨平台编译 即在windows平台下编译可以在linux平台运行的二进制文件

	SET CGO_ENABLE=0  //禁用CGO 因为跨平台编译会出现问题
	SET GOOS=linux  //设定目标操作平台
	SET GOARCH=amd64 //目标处理器架构amd64

	linux,windows,mac平台下编译跨平台设置各有不同,这里不再说明,需要的时候再去找

	go语言 package关键字 声明代码所在的包 如果为main包 最终会生成一个可执行的二进制文件
*/

/*
	go函数外面只能放(变量/常量/函数/类型)声明

	标识符与关键字:
		go语言标识符只能以 字母/数字/下划线 组成并且只能以字母或下划线开头,与其他语言一样
		关键字26个 保留字37个

*/

func main() {
	fmt.Print("hello world")            //打印不换行
	fmt.Printf("world %s hello", "123") //带占位符
	fmt.Println("hello")                //换行打印
	//这部分总体语法跟java差不多

	//变量
	fmt.Println("变量:")
	variable()

	//常量
	constant()

	//整数
	integers()

	//浮点数
	floating()

	//bool类型
	boolean()

	//字符串
	string1()
}

/*
	变量
	数据的别名
	go变量需要先声明再使用
	go 变量的声明
	全局变量可以只声明而不适用,函数内变量声明后必须使用

*/

//声明一个全局变量
var s1 string

//批量声明
var (
	string2 string //""
	int1    int    //0
	bool1   bool   //false
)

//全局变量也可以这样声明 go语言可以根据声明后赋值来推测变量的数据类型
var string3 = "s3"

func variable() {
	//变量的默认值是对应数据类型的0值
	fmt.Println(s1)
	fmt.Println(string2)
	fmt.Println(int1)
	fmt.Println(bool1)

	//变量的声明方式 当然,同一个作用域不能重复声明同名变量
	var s1 string
	var s2 = 1
	//简短声明 只能在函数内使用,不能在函数外使用
	s3 := "1"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	//匿名变量 被丢弃不用的变量通常用匿名变量接收 匿名变量不占命名空间,不会分配内存
	_ = 1
}

/*
	go语言变量函数等推荐使用驼峰式命名,这一点跟java很像
	不过go语言对于private和public函数的区分方式是看首字母大小写,首字母大写的为public函数,小写的为private函数

	go语言函数外的每一个语句都要以关键字开始 如:func var const
*/
/*
	常量的声明
*/
const pi = 3.14

//批量声明
const (
	ok = false
)

//还有这种写法 批量声明没写值默认与上一行相同
const (
	a1 = 100
	a2 //100
	a3 //100
)

//iota const关键字出现的时候iota重置为0
const (
	b1      = iota               //0
	b2                           //1
	b3                           //2
	_                            //3
	b4                           //4
	b5      = 100                //iota=5
	b6      = iota               //6
	b7, b8  = iota + 1, iota + 2 //iota=7
	b9, b10 = iota + 1, iota + 2 //iota=8
)

//使用iota
const (
	B  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func constant() {
	fmt.Println("常量:")
	fmt.Println("ok = ", ok)
	fmt.Println("a1 = ", a1)
	fmt.Println("a2 = ", a2)
	fmt.Println("a3 = ", a3)
	fmt.Println("b1 = ", b1)
	fmt.Println("b2 = ", b2)
	fmt.Println("b3 = ", b3)
	fmt.Println("b4 = ", b4)
	fmt.Println("b5 = ", b5)
	fmt.Println("b6 = ", b6)
	fmt.Println("b7 = ", b7)
	fmt.Println("b8 = ", b8)
	fmt.Println("b9 = ", b9)
	fmt.Println("b10 = ", b10)

}

/*
	go语言的基本数据类型: 整型 浮点型 bool 字符串
	整型 :
	int8 int16 int32 int64
	无符号整型
	uint8 uint16 uint32 uint64
	特殊整型
	uint 在32位机器上就是uint32 64位机器上就是uint64
	int 同理
	uintptr 用于存放指针

	在编写跨平台应用时需要注意使用int和uint的坑

*/

func integers() {
	// go语言无法直接定义二进制数
	fmt.Println("十进制数:", 100)   //100
	fmt.Println("八进制数:", 077)   //63
	fmt.Println("十六进制数:", 0xff) //255
	//输出时进制转换
	i := 100
	fmt.Printf("十进制表示:%d\n", i)  //100
	fmt.Printf("二进制表示:%b\n", i)  //1100100
	fmt.Printf("八进制表示:%o\n", i)  //144
	fmt.Printf("十六进制表示:%x\n", i) //64
	//会用到八进制给文件设置权限, 设置内存地址用到16进制
	fmt.Printf("查看变量的类型: %T\n", i)
	//强制类型转换声明
	i8 := int8(10)
	fmt.Printf("int8类型 : %T\n", i8)
}

func floating() {
	//定义一个浮点数 默认是float64
	f1 := 1.0
	fmt.Printf("浮点数的默认类型: %T\n", f1)
	f2 := float32(1.0)
	fmt.Printf("指定浮点数类型: %T\n", f2)
	//另外 f1 = f2 不能赋值 与java语言不同 java小类型可以自动转换为大范围类型 go不支持
}

//go语言还有复数 由于实际使用并不能接触到,这里不做演示

func boolean() {
	var b bool
	fmt.Print("布尔类型默认值:", b)
	//需要注意的是 go语言整型不能强转成bool类型 同时bool类型也不能转成其他类型 没有c语言 if 0  这种操作
}

/*
	go语言字符串用"" 字符用 '' 这点与java相同
	字符的转义:
	\r 回车
	\t 制表
	\n 换行

*/
func string1() {
	s2 := "你好"
	fmt.Printf("\n打印变量的值: %v\n", s2)
	fmt.Printf("打印字符串: %s\n", s2)
	//	%#v可以区分类型
	fmt.Printf("打印值加上#号: %#v\n", s2)

	//打印多行字符串 并且格式会保留
	s3 := `
		第一行
		第二行
		`
	fmt.Printf("多行字符串: %v\n", s3)

	//字符串的一些常用操作
	fmt.Println(len(s3))
	fmt.Printf("字符串相加: %v\n", s3+s2)
	fmt.Println(strings.Split(s3, "第"))     //使用后面的字符串切割 切割的结果是 []string
	fmt.Println(strings.Contains(s3, s2))   //判断s3中是否包含s2
	fmt.Println(strings.HasPrefix(s3, "第")) //判断前缀 是不是这个字符串
	fmt.Println(strings.HasSuffix(s3, "第")) //判断后缀
	fmt.Println(strings.Index(s3, "第"))     //子串首次出现的位子
	fmt.Println(strings.LastIndex(s3, "第")) //字符串最后一次出现的位子

	s4 := []string{"hell", "o"}
	fmt.Println(strings.Join(s4, s2)) //s4中字符串用s2拼接起来

	//字符串循环打印 %c 表示以字符为单位输出 如果直接打印 print(e) 就会直接打印unit8类型字符
	//unit8 一个ASCII字符
	//rune类型 表示一个utf-8字符
	for _, e := range s3 {
		fmt.Printf("%c\n", e)
	}

	/*
		字符串不能直接修改 如果想修改字符串可以用下面的方式
	*/
	s5 := "我也挺好"
	s6 := []rune(s5)
	s6[0] = '他' //rune类型 int32
	fmt.Printf("修改前: %s \n修改后: %s\n", s5, string(s6))
	fmt.Printf("字符可以转成整数: %d\n", 'A')
}

/*
	go语言的类型转换:
	整型可以和浮点型相互强制转换
	字符和字符串可以相互转换
*/
