package main

import (
	"fmt"
	"net"
)

/*
	自建一个socket服务
	由于tcp的nagle算法会产生粘包问题,可以设定一个工具用4字节记录每次发的数据包长度,接收时同样通过工具解析,来解决tcp的粘包
*/
func main() {
	//server listen port
	listener, err := net.Listen("tcp", "127.0.0.1:12121")
	if err != nil {
		fmt.Println("creat listener failed,err:", err)
		return
	}
	//获取响应
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener get conn failed,err:", err)
		return
	}
	//读数据
	var data [1024]byte
	n, err := conn.Read(data[:])
	if err != nil {
		fmt.Println("conn read data failed,err:", err)
		return
	}
	//打印数据
	fmt.Println(string(data[:n]))

}
