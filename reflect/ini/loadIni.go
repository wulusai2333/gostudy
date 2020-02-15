package ini

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

/*
	读取ini文件配置信息
	先根据路径获取文件对象
	按行读取文件
	切割字段
	根据反射给结构体赋值
*/
/*
	name{
		key:[]value
	}
	在这个配置文件的编写过程中遇到了不少问题:
	1.在使用引用类型map的同时使用了utils.Bytes2str()函数,这个函数是改变指针非值复制,
然而,本函数使用了bufio.NewReader(*File).ReadLine() 方法,此方法返回的line为[]byte 底层的数组应该是重复使用了
这就导致了保存的string字符串实际上是指向line底层数组的指针,新读取数据会修改此部分的值,从而看似值类型的string也被修改了
(以上都是我个人的猜测,debug过程中仅发现少部分字段被未知原因修改,大部分字段还是正确的显示的)
经此,对于需要保存不再修改的数据,还是用值复制的方式,而不要过分追求性能用引用
	2.debug过程是痛苦的,我在整个代码全部实现完之后才运行测试,对于一个初学者,并且是第一次编写此类代码的人来说这真的很危险,老鸟都会翻车的说
稳定安全的编写应该是写一个功能点就测一个,避免混沌(因为变化的可能性太多导致所有的预测失效),每个功能点都测通了整合的时候翻车的概率大大降低,
我这200行不到的代码没提前测,debug都整了一晚上,对于大型程序绝对是灾难性的(即使是新增一天的工作量),不得不说debug真的好用,没有debug就自己
看代码纠错疫情退了也找不出来,说不定全盘推翻重写了(虽然现在也觉得这个代码写的蛮烂的)
	3.日志模式完全没用过,明明前面已经写好了日志类,其实感觉这也没办法,日志通常是对自己觉得可能出现问题或者需要记录的点进行标记的,我这还没整啥功能呢
养成标记日志点也是个好习惯
	4.写代码前的设计时间过少,学习前面非功能性代码留下的坏习惯,键随意动,想到哪写到哪不适合写功能,不说为了性能,单单可行性也应该需要仔细思考实验一下,
编写的最后,反射大失败,此时我还是没学会反射啊,map取值赋值给对象都做不到啊,这个map套map再套切片是怎么想出来的,不说切片可以:=初始化,map在这就得make,
更灾难的是起初我内外两层的map都忘记make了...
	5.对于ini文件的理解,我找了个ini文件作参考,注释以英文;开始,之后都是注释,因为没看到有区块注释,所以我认为ini文件不存在区块注释,
[]包裹的是一个名词,下面key=value1,value2...是键值,
所以 ',' 是value的分隔符,一个key可以对应多个value,这就引发了一个大难题,拿到value之后该如何拆分,判断字节切片中是否有 ',' 没有的保存为单个值,有的保存为数组或者切片等,
存值时又有另一个问题,key姑且可以肯定是string类型,但是value可能是任意类型,在类型转换时需要考虑反射的类型转换,而要先确定接收对象结构体字段的类型,
并对value进行相应类型的转换,这里就想到了获取字段的底层类型kind()函数kind.name就可以了
对于key与字段名的映射关系,我认为需要tag来确认,有对应tag标记的才映射
根据上面的修改line就可能导致底层数组的值改变,下面的[]byte保存的值感觉有很大风险,可能需要获取到值的同时直接反射把值复制进目标对象中
*/
var ini = make(map[string]map[string][]byte, 10)

/*
	载入配置文件
*/
func LoadIni(x interface{}, tagName string) {
	configFile, err := os.OpenFile("reflect/ini/config.ini", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	//得到一个文件读写对象
	reader := bufio.NewReader(configFile)
	//读取配置到map
	readConf(reader)
	i := ini
	fmt.Println(i)
	//set value start
	//获取类型信息
	t := reflect.TypeOf(x) //t 是目标结构体的类型
	//获取值
	v := reflect.ValueOf(x) //v 是目标结构体的值
	//TODO 赋值还未成功实现
	//将值赋予对象
	var confName = t.Name()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)                                //获取字段 f是目标结构体的一个字段
		tag := f.Tag.Get(tagName)                      //字段的标签名
		if value := ini[confName][tag]; value != nil { //该字段的值
			var values = make([][]byte, 5, 10) //字段值集合
			for k := 0; k < len(value); k++ {
				if value[k] == ',' {
					values[len(values)-1] = value[:k]
					value = value[k+1:]
					k = 0
					continue
				}
				values[len(values)-1] = value[:k]
			}
			var fieldValue string //字段的值
			for index := range values {
				fieldValue = string(values[index])
			}
			//判断字段类型名
			switch c := f.Type.Kind().String(); c {
			case "int":
				a, err := strconv.ParseInt(fieldValue, 10, 64)
				if err != nil {
					fmt.Println("change fieldValue type to int failed,err:", err)
				}
				v.Field(i).SetInt(a)
			case "string":
				if fieldValue != "" {
					//a := fieldValue
					//v.Field(i).SetString(a)
				}

			default:

			}

		}
	}
	//set value end
}

func setValue(t reflect.Type, tagName string, v reflect.Value) {

}

/*
	按行读取配置文件并存入map
*/
func readConf(reader *bufio.Reader) {
	//var line []byte
	//配置名
	var confName string //这里有问题 上层循环每次循环到这里就会为confName重置
	//读取配置文件
	for {
		line, isPrefix, err := reader.ReadLine()
		//readline start
		if err != nil {
			fmt.Println("read file end...")
			break
		}
		if len(line) != 0 {
			if isPrefix {
				panic("this line is too long")
			}

			line = cleanNoteAndEndSpace(line) //清除首尾空格和注释

			//var fieldName string
			//key
			//var fieldValues []byte
			//values
			for index := range line {

				if line[index] == ']' {
					confName = string(line[1:index])
					fmt.Println(index, "confName:", confName)
					//得到一个配置名字
					///这里不用 utils.Bytes2str(s)) 因为此方法会产生一个无法预料的错误导致此变量的指针指向一个未知的位置
					if ini[confName] == nil {
						ini[confName] = make(map[string][]byte, 10)
					}
					break
				}
				//分离字段的 key:value
				/*
					对于需要保存的值,值复制是一种安全的方法,改变指针的方式对于读内存并且要复用底层数组的
					修改前:map[Mysql:map[ip:[51 52 53 54 49 46 49] name:[114 111 111 116] password:[49 50 51 52 53 54] port:[51 51 48 54]]]
					修改后:map[Mysql:map[ip:[51 52 53 54 49 46 49] name:[114 111 111 116] password:[49 50 51 52 53 54] port:[51 51 48 54]]]
				*/
				if line[index] == '=' {
					fieldName := string(line[:index])
					va := string(line[index+1:])
					fmt.Println(fieldName, va) //打印字段名
					fieldValues := line[index+1:]
					fmt.Println(fieldValues) //打印字段值
					//根据字段名将value存入
					if confName != "" {
						fmt.Println("confName:", confName)
						ini[confName][fieldName] = fieldValues
						fmt.Println("字段:", ini[confName][fieldName])
					}
					break
				}

			}

		}

		//readline end
	}
	//sline, err := reader.ReadString('\n')

}

//去掉注释和行末空格
func cleanNoteAndEndSpace(line []byte) []byte {
	//先将注释去掉
	for index := range line {
		if line[index] == ';' {
			line = line[:index]
			break
		}
	}
	//去掉行末空格
	for index := len(line) - 1; index >= 0; index-- {
		if line[index] != ' ' {
			//line = line[:len(line)-index]
		}
	}
	return line
}
