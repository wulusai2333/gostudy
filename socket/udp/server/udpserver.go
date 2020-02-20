package main

import (
	"fmt"
	"net"
	"strings"
)

/*
	udp服务端,接收客户端信息
*/
func main() {
	//建立客户端连接
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("create udp conn failed,err:", err)
		return
	}
	defer conn.Close()
	var data [1024]byte
	//读取客户端发送数据

	i, addr, err := conn.ReadFromUDP(data[:])
	if err != nil {
		fmt.Println("server conn read data failed,err:", err)
	}
	fmt.Println(addr.String(), data[:i])
	//返回数据
	replace := strings.ToUpper(string(data[:i]))
	n, err := conn.WriteToUDP([]byte(replace), addr)
	if err != nil {
		fmt.Println("send replace to udp failed,err:", err)
		return
	}
	fmt.Println(n)

}
