package main

import (
	"fmt"
	"strings"
)

/*
	panic 相当于java抛出了严重异常,程序直接停了
	recover 相当于java的try-catch接收了异常,程序跳过异常部分接着执行
	注意: recover必须搭配defer使用 defer一定要在能引发panic的地方之前定义
*/
func main() {
	funA()
	funB()
	funC()
	_printf()
	_scan()
	_coin()
}

func funA() {
	fmt.Println("A...")
}
func funB() {
	defer func() {
		fmt.Println("关闭链接...")
		err := recover()
		fmt.Println(err)
	}()
	panic("error!")
	fmt.Println("B...")
}
func funC() {
	fmt.Println("C...")
}

/*
	fmt.Printf()打印
	Sprintf 会返回打印的值为一个字符串
	scan 扫描
	scanf 可以用下面格式化扫描
	fscan 扫描文件
	%T	查看类型
	%d	十进制
	%b	二进制
	%o	八进制
	%x	十六进制a-f
	%X	十六进制A-F
	%c	字符
	%s	字符串
	%p	指针
	%v	值
	%f	浮点数
	%t  布尔
	%U  unicode表示	U+1234
	%q	go语法字符字面值	整数->字符
	%e %E 科学计数法	1.0E+12
	%g %G 根据实际情况转换为%e或%f
	-----格式化占位符
	%v 通用的
	%+v 输出结构体时会添加字段名
	%#v 值的go语法表示
	%T	值类型
	%% 百分号
*/
func _printf() {
	s := "wlc"
	fmt.Printf("%%T :%T \n", s)
	fmt.Printf("%%v :%v \n", s)
	fmt.Printf("%%+v :%+v \n", s)
	fmt.Printf("%%#v :%#v \n", s)
	fmt.Printf("%% :%% \n")
	fmt.Printf("%%U :%U \n", 's')
	fmt.Printf("%%q :%q \n", 65)
}

/*
	课程名 章节序号 章节名 章节链接
	如果分表的话	*sign标记显示隐藏
	课程序号	课程名 sign
	章节序号	章节名 课程序号 章节链接 sign
*/
/*
	把用户输入内容存到s的地址位置
	修改某个变量值传指针
*/
func _scan() {
	var s string
	fmt.Scan(&s)
	fmt.Println(s)
}

/*
	分金币 根据每个人名字里出现的字母分不同的金币,计算每个人分的金币和最后剩余
	e E 每出现一次+1金币
	i I	每出现一次+2金币
	o O 每出现一次+3金币
	u U 每出现一次+4金币

*/
func _coin() {
	coins := 1000
	users := []string{"dasef", "egesgfb", "wafeagir", "watweuafds", "rhtyfoggr", "weagrgisds", "asdfg"}
	fen := make(map[string]int, len(users))
	m, coin := fenCoin(coins, users, fen)
	fmt.Println(m, coin)
}
func fenCoin(coins int, users []string, fen map[string]int) (map[string]int, int) {

	for _, name := range users {
		for _, c := range strings.Split(strings.ToUpper(name), "") {
			switch c {
			case "E":
				fen[name] += 1
				coins -= 1
			case "I":
				fen[name] += 2
				coins -= 2
			case "O":
				fen[name] += 3
				coins -= 3
			case "U":
				fen[name] += 4
				coins -= 4
			default:
				fen[name] += 0
			}
		}
	}
	return fen, coins

}
