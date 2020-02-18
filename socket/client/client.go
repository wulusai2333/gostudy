package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12121")
	if err != nil {
		fmt.Println("client get conn failed,err:", err)
		return
	}
	var data = []byte("hello")
	n, err := conn.Write(data)
	if err != nil {
		fmt.Println("conn write data failed,err:", err)
		return
	}
	fmt.Println(n)
}
