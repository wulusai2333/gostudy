package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
	http 客户端发送get请求到服务端,接收服务端返回的数据
*/
//长连接 连接复用
var (
	client = http.Client{
		Transport: &http.Transport{DisableKeepAlives: false},
	}
)

func main() {
	//http.get发送get请求
	http_get()
	//http.NewRequest指定请求方式发送请求
	http_new_request()
	//http_post()
	//http_post_form()
}

func http_post_form() {
	response, err := http.PostForm("http://127.0.0.1:10086/upload", url.Values{
		"key": {"value"}, "id": {"11"},
	})
	if err != nil {
		fmt.Println("post form failed,err：", err)
		return
	}
	defer response.Body.Close()
}

func http_post() {
	buf := bufio.Reader{}
	resp, err := http.Post("http://127.0.0.1:10086/upload", "image/jpeg", &buf)
	if err != nil {
		fmt.Println("post failed,err:", err)
		return
	}
	defer resp.Body.Close()
}

func http_new_request() {
	//?后面参数
	values := url.Values{}
	values.Set("name", "xiaohong")
	values.Set("age", "11")
	encode := values.Encode()
	//设置连接属性
	urlStr, err := url.ParseRequestURI("http://127.0.0.1:10086/user")
	if err != nil {
		fmt.Println("url parse failed,err:", err)
		return
	}
	urlStr.RawQuery = encode
	//获得请求对象
	request, err := http.NewRequest("GET", urlStr.String(), nil)
	if err != nil {
		fmt.Println("client new request send failed,err:", err)
		return
	}
	//发送请求获得相应对象
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("client do request failed,err:", err)
		return
	}

	defer response.Body.Close() //关闭连接
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

	//禁用keep alive 的短连接
	tr := &http.Transport{DisableKeepAlives: true}
	client := http.Client{Transport: tr}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("client do failed,err:", err)
		return
	}
	defer resp.Body.Close()
}

//使用http.get发送get请求
func http_get() {
	resp, err := http.Get("http://127.0.0.1:10086/user?name=xiaoming&age=12")
	if err != nil {
		fmt.Println("http get failed,err:", err)
		return
	}
	defer resp.Body.Close() //记得关闭连接
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("client read resp body err", err)
		return
	}
	fmt.Println(string(bytes))
}
