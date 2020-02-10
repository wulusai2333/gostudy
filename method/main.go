package main

import "fmt"

/*
	方法作用特定函数
*/
type dog struct {
	name string
	age  int
}

func newDog(name string, age int) *dog {
	return &dog{name: name, age: age}
}

/*
	前面限定了dog类型,这个函数只能dog类型参数能调用
*/
func (d dog) wang() {
	fmt.Printf("%v:汪汪汪~\n", d.name)
}

/*
	前面的类型传的值,所以他本身并没有被改变
*/
func (d dog) laole() {
	d.age++
}

/*
	传递的是指针,所以age被真正改变了
*/
func (d *dog) zhenlaole() {
	d.age++
}
func main() {
	d := newDog("小黄", 11)
	d.wang()
	fmt.Println(d.age)
	d.laole()
	fmt.Println(d.age)
	d.zhenlaole()
	fmt.Println(d.age)
	//是用自定义类型的方法
	// myint(100)是什么,就是对int类型强转
	i := myint(100)
	i.heihei()
}

/*
	给自定义类型添加方法
*/
type myint int

func (m myint) heihei() {
	fmt.Println("heihei ", m)
}
