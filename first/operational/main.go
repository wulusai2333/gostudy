package main

import (
	"fmt"
	"math/rand"
)

/*
	go语言的运算符
	+ - * / % 算数运算符 逻辑与其他语言没啥区别,就不做演示
	注意 ++ -- 在go语言中作为单独语句,不是运算符,不能放在等号右边赋值 也就没有 a=a++ 这种操作
	== != >= <= < > 比较运算符 结果为true或者false
	&& || ! 逻辑运算符
	& | ! << >> 位运算符
	= += -= *= /= %= <<= >>= &= |= ^= 赋值运算符 a+=b ==> a = a+b
	这部分在java中学的够多了,go语言在这里也没多少区别,此处不做演示
*/

/*
	复合数据类型:
	数组 需指定长度和数据类型
	var arr [2]int
	切片 不需要指定,自动扩容
	var s []int
	map 键值对存储,无序
	var m map[int]string
	切片和map使用前都需要初始化 都是引用类型,数组是值类型
	值类型和引用类型的区别:值类型声明时都是对应数据类型的0值,引用类型声明后是nil,初始化后才会在内存中开辟空间
	make([]string,10,20)
*/

func main() {
	_array()
	_slice()
	_ptr()
	_map()
}

func _array() {
	//普通的定义一个数组
	var arr = [3]int{1, 2, 5}
	for k, v := range arr {
		fmt.Printf("k: %d v:%d\n", k, v)
	}
	//数组钱定义的长度可以省略
	var arr2 = [...]string{"w", "b", "z", "d"}
	fmt.Println(arr2)
	//二维数组 外层可以省略,但是内层不能省略...
	var arr3 = [...][3]int{
		[...]int{1, 2, 3},
		{4, 5, 6}, //需要注意的末尾的 ,
	}
	fmt.Println(arr3)

	/*
		go语言 = 结果是拷贝值 而非赋值地址 所以有如下
		数组是值类型
	*/
	arr4 := arr
	arr4[1] = 3
	fmt.Println(arr)  //[1 2 5]
	fmt.Println(arr4) //[1 3 5]

	//求数组值的和
	var sum int
	for _, v := range arr {
		sum += v
	}
	fmt.Println("数组的和为:", sum)
}

/*
	与数组不同的是切片指定的是 底层数组 开始索引 切片长度
	切片未指定长度,切片随着存入东西的增加可以自动扩容,小于1024长度扩容直接翻倍,大于1024,newcap+=newcap/4,
*/
func _slice() {
	//切片是引用类型不能直接比较,只能和nil比较
	var s1 []string
	var s2 []int
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//切片的初始化方式
	arr := [7]string{"a", "w", "s", "l", "t", "q", "l"}
	s1 = arr[1:3] //从索引1到索引3含头不含尾 [w s]
	s3 := arr[:4] // [0,4) [a w s l]
	s4 := arr[3:] //	[3,len(arr)) [l t q l]
	s5 := arr[:]  //从开始切到最后 [a w s l t q l]
	fmt.Println(s1)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	s2 = []int{1, 2, 3}
	//len求切片长度 cap求切片容量,与底层数组容量有关
	fmt.Println(len(s1)) //2
	fmt.Println(cap(s1)) //6

	//切片再切片
	s6 := s4[1:]                                                         //[t q l]
	fmt.Printf("s6: %v ,len(s6):%v ,cap(s6):%v\n", s6, len(s6), cap(s6)) //s6: [t q l] ,len(s6):3 ,cap(s6):3

	//修改底层数组值 切片值肯定会变
	arr[5] = "m"
	fmt.Println(s4) //[l t m l]
	fmt.Println(s6) //[t m l]

	//  make([]T,size,cap) 动态创建切片,不需要依赖数组来创建 cap可以不指定 默认等于size
	s7 := make([]int, 10, 20)
	fmt.Printf("s7: %v ,len(s7):%v ,cap(s7):%v\n", s7, len(s7), cap(s7))

	//切片相当于java 的List 里面保存的是地址 而切片不同的是只能保存连续的内存(实质上只能存一个地址,而List可以保存多个地址,甚至不同数据类型的地址)

	//判断切片是否为空 而不用s7 == nil判断 nil代表没有对应底层数组 而len==0对应了nil的情况和切片长度为0的情况,这个需要根据实际情况来判断如何使用
	fmt.Println(len(s7))

	//修改切片的值 实际上修改的是底层数组的值
	s4[2] = "p"
	fmt.Println(s4)
	fmt.Println(s6)
	fmt.Println(arr)

	//切片的遍历 跟数组一样

	for k, v := range s6 {
		fmt.Println(k, " ", v)
	}

	//删除切片中元素 实质上相当于修改了底层数组的元素,底层数组还是那个底层数组
	fmt.Println("s4删除元素前:", s4)
	s4 = append(s4[0:1], s4[2:]...)
	fmt.Println("s4删除元素后:", s4)
	fmt.Println("底层数组:", arr)
	s4 = append(s4[0:1], "o")
	fmt.Println("底层数组:", arr)
}

/*
	指针
	go语言不存在指针操作
	& 取地址
	* 根据地址取值
*/
func _ptr() {
	a := 0
	p := &a
	fmt.Println("p的值:", p)
	fmt.Printf("p的类型:%T\n", p)

	//这个是有问题的 因为 b=nil *b没有对应的地址,也就无从谈给值赋值了
	var b *int
	b = p //给b一个地址就可以了
	*b = 100
	fmt.Println(*b)

	//new的使用
	var c *int
	c = new(int)
	*c = 1
	fmt.Println(*c)
	/*
		make 与 new 的区别
		make只用于 切片 map chan的创建 返回的是这三个类型本身而不是内存地址
		new 创建返回的是类型的地址,给基本数据类型申请地址的
		func new(Type) *Type
	*/
}

/*
	map 无序的key-value数据结构 是引用类型
*/
func _map() {
	var m map[int]string
	fmt.Println(m == nil)       //没有初始化 不能直接用 没有在内存中开辟空间
	m = make(map[int]string, 5) //初始化时定好大小,避免在运行期间再动态扩容
	m[1] = "zzz"
	m[2] = "aaa"

	fmt.Println(m)
	//取值
	s, ok := m[1]
	if ok {
		fmt.Println(s)
	}
	//遍历 map是以哈希的方式存储key,存的值与存入的顺序无关
	//如果只想取key  k:=range m
	//只想取value 	_,v:=range m
	for k, v := range m {
		fmt.Printf("key:%d value:%s\n", k, v)
	}
	/*
		如果想把map存入的数据排序取出来
			先将key取出排序 再根据排好顺序的key取值
	*/
	m1 := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		m1[rand.Intn(100)] = i
	}
	fmt.Printf("未排序map:%v\n", m1)
	s1 := make([]int, 10, 10)
	i := 0
	for k := range m1 {
		s1[i] = k
		i++
	}
	fmt.Printf("排序好的key:%v\n", s1)
	/*
		切片和map的嵌套
		需要注意的是 切片和map都需要make开辟空间
	*/
	s2 := make([]map[int]int, 10, 10)
	s2[0] = make(map[int]int)
	s2[0][100] = 10
	fmt.Printf("初始化完成的切片存map:%v\n", s2)

	m2 := make(map[int][]int, 10)
	m2[10] = []int{10, 12, 23}
	fmt.Println(m2)
	fmt.Println(m2[1]) //key对应的value不存在,返回的是对应value类型的0值
	delete(m2, 1)      //删除对应的key
}
