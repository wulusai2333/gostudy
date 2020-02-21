package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof" //只需访问 http://127.0.0.1:10086/debug/pprof/ 即可查看信息
	"os"
)

/*
	原生http服务器
	安装graphviz画内存图 和 perl语言生成火焰图
安装 安装FlameGraph脚本
git clone https://github.com/brendangregg/FlameGraph.git
cp FlameGraph/flamegraph.pl /usr/local/bin #linux服务器中操作
现在已经可以不用 go-torch 来生成火焰图了，安装 pprof
go get -u github.com/google/pprof
然后使用
pprof -http 127.0.0.1:9090 http://127.0.0.1:8080/debug/pprof/profile 就可在127.0.0.1:9090地址看火焰图了
*/
func main() {
	// 主函数中添加
	go func() {
		http.HandleFunc("/program/html", index) // 用来查看自定义的内容
		log.Println(http.ListenAndServe("0.0.0.0:8080", nil))
	}()
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
	file, err := os.Open("./static/index.html")
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
