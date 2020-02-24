package main

import (
	"bufio"
	"fmt"
	"github.com/wulusai2333/gostudy/search/search"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	_, err := os.Stat("./filepath.txt")
	if os.IsNotExist(err) {
		go search.LoadData()
	}
	http.HandleFunc("/index", index)
	http.HandleFunc("/find", find)
	http.HandleFunc("/js/vue.js", js)
	http.HandleFunc("/js/axios-0.18.0.js", js)
	err = http.ListenAndServe("127.0.0.1:12345", nil)
	if err != nil {
		fmt.Println("http server start failed,err:", err)
	}
}

func js(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
	//fmt.Println(request.URL.Host)
	file, err := os.Open("./static" + request.URL.Path)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	var data [1024]byte
	for {
		n, err := file.Read(data[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read file failed,err:", err)
			return
		}
		_, err = writer.Write(data[:n])
		if err != nil {
			fmt.Println("response write failed,err:", err)
			return
		}
	}

}

var files = make(map[string][]byte, 1000)

/*
	这里应该做成 获取文件名,查找文件名中是否包含字段,而不是全路径是否包含字段
*/
func find(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	fmt.Println(values.Get("name"))
	name := values.Get("name")
	if name == "" {
		_, err := writer.Write([]byte("null"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if len(files) == 0 {
		file, err := os.OpenFile("./filepath.txt", os.O_RDONLY, 0644)
		if err != nil {
			fmt.Println("open file failed,err:", err)
			return
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			value := append([]byte{}, line...)
			fileSplit := strings.Split(string(line), "\\")
			filename := fileSplit[len(fileSplit)-1]
			files[filename] = value

		}
	}

	// if end
	for k, v := range files {
		if strings.Contains(k, name) {
			//fmt.Println(k, string(v))
			_, err := writer.Write(v)
			if err != nil {
				fmt.Println("response write failed,err:", err)
				return
			}
		}
	}

}

func index(writer http.ResponseWriter, request *http.Request) {
	file, err := os.Open("./static/index.html")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	var data [1024]byte
	n, err := file.Read(data[:])
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	_, err = writer.Write(data[:n])
	if err != nil {
		fmt.Println("response write failed,err:", err)
		return
	}
}
