package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*

	打开文件
	*File,err := os.Open("file/index")
	defer os.close()

	文件读写的几种方法
*/
func main() {
	readFile()
	read_use_bufio()
	read_user_ioutil()
	file_write()
	bufio_write()
	ioutil_write()
	//use_bufio_get_user_input() //获取用户输入 可以带空格 这里注释掉以免影响程序运行流畅性
	copyFile()
	insertFile()
}

/*
	关闭文件资源
*/
/*func fileClose(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Println("file close failed,err", err)
		return
	}
}*/
/*
	在文件中间插入字符
	具体思路:
		将文件指针移动到光标处

*/
func insertFile() {
	file, err := os.OpenFile("log/a.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed,err: %v\n", err)
		return
	}
	defer file.Close()
	//将文件指针移动到3字节处
	_, err = file.Seek(3, 0)
	if err != nil {
		fmt.Println("seek file failed,err:", err)
	}
	//覆盖修改后面的字节
	_, err = file.Write([]byte{'a', 'b', 'c'})
	if err != nil {
		fmt.Println("write file failed, err:", err)
	}
}

/*
	复制文件
*/
func copyFile() {
	/*
		copy from
	*/
	fileRead, err := os.Open("package/main.go")
	if err != nil {
		fmt.Printf("open file failed,err: %v\n", err)
		return
	}
	defer fileRead.Close()
	/*
		copy to
	*/
	fileWrite, err := os.OpenFile("log/a.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer fileWrite.Close()

	/*
		copy
	*/
	var data [1024]byte //创建一个字节数组
	for {
		n, err := fileRead.Read(data[:]) //读取字节
		if err == io.EOF {
			fmt.Println("文件已经读完了!")
			return
		}
		if err != nil {
			fmt.Printf("read file failed,err: %v\n", err)
			return
		}
		_, err = fileWrite.Write(data[:n]) //写入字节
		if err != nil {
			fmt.Println("write file failed, err:", err)
			return
		}
	}
}

func use_bufio_get_user_input() {
	//获取标准输入 遇到如下字符才停止
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read Stdin failed, err:", err)
		return
	}
	fmt.Println(s)
}

func ioutil_write() {
	str := "全自动烤鸭切片\n"
	//这个方法会覆盖写入文件
	err := ioutil.WriteFile("log/a.txt", []byte(str), 0644)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func bufio_write() {
	file, err := os.OpenFile("log/a.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	wr := bufio.NewWriter(file)
	_, err = wr.WriteString("我想吃烤鸡!\n")
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
	err = wr.Flush()
	if err != nil {
		fmt.Println("flush file failed, err:", err)
		return
	}
}

func file_write() {
	file, err := os.OpenFile("log/a.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello world!\n"
	_, err = file.Write([]byte(str))
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
	_, err = file.WriteString("你好!\n")
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

/*
	ioutil包
	直接读取文件
*/
func read_user_ioutil() {
	file, err := ioutil.ReadFile("interface/main.go")
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	fmt.Print(string(file))
}

/*
	bufio读取
	提供了丰富的读取方式,下面展示的是按行读取
*/
func read_use_bufio() {
	file, err := os.Open("log/main.go")
	if err != nil {
		fmt.Printf("open file failed,err: %v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读完了!")
			return
		}
		if err != nil {
			fmt.Println("read file failed,err:", err)
			return
		}
		fmt.Print(line)
	}
}

/*
	字节读取
*/
func readFile() {
	file, err := os.Open("package/main.go")
	if err != nil {
		fmt.Printf("open file failed,err: %v\n", err)
		return
	}
	defer file.Close()
	//读文件
	var data [128]byte
	for {
		n, err := file.Read(data[:])
		if err == io.EOF {
			fmt.Println("文件已经读完了!")
			return
		}
		if err != nil {
			fmt.Printf("read file failed,err: %v\n", err)
			return
		}

		fmt.Printf("读取到%d个字节\n", n)
		fmt.Print(string(data[:n]))
		if n < 128 {
			return
		}
	}
}
