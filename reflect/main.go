package main

import (
	"fmt"
	"github.com/wulusai2333/gostudy/reflect/ini"
	"reflect"
)

/*
	reflect 反射
*/
func reflectType(x interface{}) {
	//反射获取字段的类型
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v type name:%v type kind:%v\n", t, t.Name(), t.Kind())
}

type person struct {
	Name string `json:"name"`
	Age  int
}

/*
  mysql 配置结构体
*/
type Mysql struct {
	Ip       string `ini:"ip"`
	Port     int64  `ini:"port"`
	Name     string `ini:"name"`
	Password string `ini:"password"`
}

func main() {
	//反射练习
	//reflectTest()
	// 这段要用命令行编译运行
	mysql := Mysql{}
	ini.LoadIni(&mysql, "ini")
	fmt.Printf("%#v", mysql)
}

//反射的练习
func reflectTest() {
	var a = 10
	reflectType(a)
	str := "100"
	reflectType(str)
	p := person{"小呜", 12}
	reflectType(p)
	//反射获取结构体内部字段信息
	getField(p)
	//通过反射修改字段的值
	b := 10
	replaceValue(&b)
	fmt.Println(b)
}

func replaceValue(b *int) {
	v := reflect.ValueOf(*b)
	//传的是b的值
	fmt.Println(v.Type())
	//v.SetInt(100)//修改的是副本 reflect包会引发panic
	fmt.Println(b)
	//想要修改值 必须得传递指针
	v = reflect.ValueOf(b)
	v.Elem().SetInt(100)
	//用指针获取值再修改
	fmt.Println(*b)
}

func getField(p person) {
	t := reflect.TypeOf(p)
	//获取类型
	//循环获取结构体所有字段信息
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("field name:%v field type:%v field tag:%v \n", f.Name, f.Type, f.Tag.Get("json"))
	}
	//通过字段名获取指定结构体字段信息
	if nameField, ok := t.FieldByName("Name"); ok {

		fmt.Printf("field name:%v field type:%v field tag:%v field index:%v\n", nameField.Name, nameField.Type, nameField.Tag.Get("json"), nameField.Index)
	}
}
