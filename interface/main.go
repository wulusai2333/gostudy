package main

import "fmt"

/*
	接口的使用
	接口只关心方法
	一个结构体可以实现多个接口 接口直接也可以套娃
	interface{} 空接口
	根据接口定义,只要实现了接口所有方法就算实现了该接口,空接口没有方法,所以任意类型都实现了该接口 (有点像java里的Object类)
	即所有类型都可以传
*/
/*
	这个就有点java的Map类的味道了
*/
type list struct {
	data map[interface{}]interface{}
}

type dog struct {
	name string
}

type cat struct {
	name string
}
type person struct {
	name string
}

func (c *cat) speak() {
	fmt.Println("喵~")
}
func (d *dog) speak() {
	fmt.Println("汪~")
}
func (p *person) speak() {
	fmt.Println("啊♂~")
}

/*
	这里都使用了指针接收者来实现
	如果使用值接收者实现 则接口既可以接收指针,也可以接收值
*/
func (p person) eat(food string) {
	fmt.Printf("%s正在吃%s\n", p.name, food)
}

/*
	定义了一个接口 接口里面放方法
	需要实现接口的全部方法,并且方法的参数也要完全一致,才算实现了这个接口
*/
type speaker interface {
	speak()
}
type animal interface {
	//speak() //实现了这个方法表示也实现了上面的接口
	//speaker 也可以直接套其他接口
	eat(string)
}

/*
	接收接口参数
*/
func dapipi(i speaker) {
	i.speak()
}
func main() {
	_interface_test()

	assign("1")
}

/*
	接口类型的断言
*/
func assign(i interface{}) {
	//断言类型,如果 ok=true则断言正确 str是类型i的值
	str, ok := i.(string)
	if ok {
		fmt.Println(str)
	}
	switch t := i.(type) {
	case string:
		fmt.Printf("this a 字符串 %v\n", t)
	case int:
		fmt.Printf("this a 数字 %v\n", t)
	case []interface{}:
		fmt.Printf("this a 切片 %v\n", t)
	case [10]interface{}:
		fmt.Printf("this a 数组 %v\n", t)
	case map[interface{}]interface{}:
		fmt.Printf("this a map %v\n", t)
	case bool:
		fmt.Printf("this a bool类型 %v\n", t)
	case interface{}:
		fmt.Printf("this a %v 我肯定猜到了\n", t)
	default:
		fmt.Println("这不可能!我居然没有猜到")
	}
}
func _interface_test() {
	//传入一个接收指定接口类型的函数 类型存在,但值为nil
	var c *cat
	//*main.cat <nil>
	var d *dog
	var p *person
	dapipi(c)
	dapipi(d)
	dapipi(p)
	//定义接口类型并赋值
	var s speaker
	s = c
	s = d
	s = p
	s.speak()
	fmt.Printf("%T\n", s)
	/*
		接口存储时分两部分  interface{struct: *main.person,value: person{}}
		初始时 interface(nil,nil) 动态类型 动态值
	*/
	var a animal
	p = &person{"人"}
	//这里如果不赋值,下面调用eat方法时就会出现错误 panic: runtime error: invalid memory address or nil pointer dereference
	a = p
	a.eat("鸡腿")
	fmt.Printf("%T %v\n", c, c)
	var b person
	// main.person {}
	fmt.Printf("%T %v\n", b, b)
	a = b
	//这里可以看出 接口时引用类型 只存储地址
	a.eat("大米")
}
