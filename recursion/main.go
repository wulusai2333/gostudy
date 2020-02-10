package main

import (
	"encoding/json"
	"fmt"
)

/*
	struct 结构体
	type 类型
	结构体 与java的类作用和用法非常相似,只是与文件名的关联性不强
	匿名结构体
	结构体是值类型,也就是定义的时候就在内存中开辟好了空间,不需要用如make去开辟空间,不能跟nil比较
*/
type myint int //自定义类型	显示的是 main.myint 类型
type me = int  //类型别名,只在当前生效,实际还是 int 类型
/*
	结构体
*/
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

/*
	构造函数使用如下方式定义
	go语言函数是值拷贝,所以如果返回的是值不是指针在大量使用时就会极大浪费系统内存
	所以下方式创建内存,返回的指针
*/
func newPerson(name string, age int, gender string, hobby []string) *person {
	return &person{
		name:   name,
		age:    age,
		gender: gender,
		hobby:  hobby,
	}
}
func fp(x person) {
	x.age = 20
}
func fp2(x *person) {
	//(*x).age = 20 //原本应该这么操作修改值的
	x.age = 20 //go的语法糖,因为指针是无法修改的,所以这样就直接转为对指针指向的内存进行操作
}
func main() {
	//类型
	_type()
	//结构体
	_struct()
	/*
		匿名结构体
	*/
	nimingstruct()
	/*
		创建一个指针类型结构体
	*/
	pstruct()
	/*
		key-value形式初始化struct
	*/
	k_v_struct()
	/*
		值列表初始化
	*/
	_value_struct()
	//结构体内存是连续的
	_neicun()
	//结构体构造函数
	use_gouza()
	//结构体的匿名字段
	niming()
	//结构体的嵌套
	_niming_qiantao_struct()
	/*
		go语言本身没有继承,但是使用上面的特性,可以实现"继承"的效果
	*/
	_jicheng()
	/*
		结构体与json 序列化与反序列化
	*/
	d := dog{
		feet:   4,
		animal: animal{},
	}
	data, err := json.Marshal(d)
	if err != nil || len(data) == 0 {
		fmt.Println("序列化失败", err)
	}
	fmt.Println(string(data)) //{}
	/*
		这里显示不了字段的值 因为字段的可见性 与函数和结构体一样,只有首字母大写的字段才能被外界访问 json这个包相对于main包就是外界了
		所以需要如下方式来创建结构体
	*/
	c := cat{
		Feet: 4,
		Name: "喵喵",
	}
	//序列化
	data, err = json.Marshal(c)
	if err != nil || len(data) == 0 {
		fmt.Println("序列化失败", err)
	}
	fmt.Println(string(data)) //{"Feet":4,"Name":"喵喵"}
	/*
		让字段名在json格式下显示为小写就需要 `json:"name"`这种注释了
	*/
	var c2 cat
	//反序列化
	if err == json.Unmarshal(data, &c2) && err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println(c2)
}

type cat struct {
	Feet uint8  `json:"feet" db:"FEET"`
	Name string `json:"name"`
}

func _jicheng() {
	d := dog{
		feet: 4,
		animal: animal{
			"大黄",
			home{
				"xinjiang",
				"2.2.2.2", //为何大括号前面也要跟 , 因为在go中换行表示语句的结束 ,告诉了编译器还没结束
			},
		},
	}
	d.wang()
	d.move()
	//可以直接调用匿名结构体的方法,有点类似java继承的结果,可以直接调用父类的方法,如果子类和父类重名,就等于重写了父类的方法(感觉go的实现方法有点妙)
	fmt.Println(d.city)
	//多层匿名结构体的字段依旧可以使用 套娃
}

type dog struct {
	feet uint8
	animal
}

func (d *dog) wang() {
	fmt.Printf("%s在叫!\n", d.name) //这里dog并没有定义name字段,却可以使用name字段,符合上面的结构体嵌套中匿名结构体的规则
}
func (a *animal) move() {
	fmt.Printf("%s在走动~\n", a.name)
}
func _niming_qiantao_struct() {
	a := animal{
		name: "zhu",
		home: home{"beij", "1.1.1.1"},
	}
	z := zoo{
		name:    "xiaomao",
		home:    home{"beiji", "1.1.2.1"},
		animals: nil,
	}
	z.animals = append(z.animals, a)
	fmt.Println(z.city)
	//匿名嵌套结构体,可以直接跳一层级取下一层级的字段 本来应该是z.home.city
	//如果字段冲突就没办法这么写了,必须像下面这样写全
	fmt.Println(z.home.city)
	fmt.Println(z.animals[0].home.ip)
}

type animal struct {
	name string
	home
}
type home struct {
	city string
	ip   string
}
type zoo struct {
	name string
	home
	animals []animal
}

func niming() {
	// 类型名即字段名,同类型名不能有两个匿名字段,要不就冲突了
	var p = struct {
		string
		int
	}{}
	p.int = 10
	p.string = "小明"
	fmt.Println(p)
}

func use_gouza() {
	p := newPerson("小呜", 18, "女", []string{"直播"})
	fmt.Printf("%v\n", *p)
}

func _neicun() {
	var p = person{
		name:   "呐",
		age:    6,
		gender: "futa",
		hobby:  []string{"女装"},
	}
	fmt.Printf("p变量内部内存连续:%p %p %p \n", &(p.name), &(p.age), &(p.gender))
}

func _value_struct() {
	//&直接拿指针
	var p = &person{
		"小呜",
		14,
		"男",
		[]string{"打电动"},
	}
	fmt.Printf("值列表初始化:%T\n %p\n %v\n", p, p, p)
}

func k_v_struct() {
	var p = person{
		name:   "小呜",
		age:    12,
		gender: "女",
		hobby:  []string{"吃饭", "睡觉", "打豆豆"},
	}
	fmt.Printf(" 声明时赋值:%T\n 变量值:%#v\n 地址:%p\n", p, p, &p)
}

func pstruct() {
	var p2 = new(person)
	fmt.Printf("指针类型结构体:%T\n", p2)
	fmt.Printf("结构体的指针:%p\n", p2)
	//下面两个帮助理解内存地址和值的关系
	var a = 10
	b := &a
	fmt.Printf("a的值:%v \na的地址:%p\n", a, &a)
	fmt.Printf("b的值:%v \nb的地址:%p\n", b, &b)
}

func _type() {
	fmt.Println(f(5))
	fmt.Println(taijie(10))
	var a myint = 100
	var b me = 100
	fmt.Printf("myint:%v,%T\n", a, a)
	fmt.Printf("me:%v,%T\n", b, b)
}

func _struct() {
	var p person
	//fmt.Println(p==nil) 不能跟nil比较
	p.name = "小呜"
	p.age = 18
	p.gender = "男"
	p.hobby = []string{"抽烟", "喝酒", "烫头"}
	fmt.Println(p)
	fp(p)
	fmt.Println(p.age)
	//age=18 在次强调 go语言函数传递的永远是拷贝
	fp2(&p)
	//传递内存地址
	fmt.Println(p.age)
	//这里就是根据内存地址修改后的值了
}

/*
	匿名结构体
*/
func nimingstruct() {
	var s struct {
		a string
		b int
	}
	s.b = 100
	s.a = "jojo"
	fmt.Println(s)
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
