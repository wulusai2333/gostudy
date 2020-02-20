package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*
	原生http服务器
*/
func main() {
	//这个写在listen前面,写在后面没用
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/index", index)
	http.HandleFunc("/user", user)
	err := http.ListenAndServe("127.0.0.1:10086", nil) //监听和服务
	if err != nil {
		fmt.Println("http server start failed,err:", err)
	}

}

//接收客户端带参数的get请求,并返回ok
func user(writer http.ResponseWriter, request *http.Request) {
	//获取url上的参数
	values := request.URL.Query()
	name := values.Get("name")
	age := values.Get("age")
	fmt.Println(name, age)
	//读请求体
	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("http server read request body failed,err:", err)
		return
	}
	fmt.Println(string(bytes))
	//写回复
	_, err = writer.Write([]byte("ok"))
	if err != nil {
		fmt.Println("writer failed,err:", err)
		return
	}
}

//首页
func index(writer http.ResponseWriter, request *http.Request) {
	file, err := os.Open("socket/http/server/static/index.html")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	var data [1024]byte
	n, err := file.Read(data[:])
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	i, err := writer.Write(data[:n])
	if err != nil {
		fmt.Println("response write failed,err:", err)
		return
	}
	fmt.Println(i)
}

//一个简单的http响应
func hello(responseWriter http.ResponseWriter, request *http.Request) {
	i, err := responseWriter.Write([]byte("hello~"))
	if err != nil {
		fmt.Println("response write failed,err:", err)
		return
	}
	fmt.Println(i)
}
