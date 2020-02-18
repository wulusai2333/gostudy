package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//要连接的服务端地址
	udpToAddr := &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	}
	/*	udpForAddr := &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 10010,
	}*/
	//建立连接
	client, err := net.DialUDP("udp", nil, udpToAddr)
	if err != nil {
		fmt.Println("create udp conn failed,err:", err)
		return
	}
	defer client.Close()
	//发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string err,err:", err)
		}
		i, err := client.WriteToUDP([]byte(s), udpToAddr)
		if err != nil {
			fmt.Println("client conn write data failed,err:", err)
		}
		fmt.Println(i)
		//接收响应数据
		var data [1024]byte
		i, addr, err := client.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("server conn read data failed,err:", err)
		}
		fmt.Println(addr.String(), data[:i])
	}

}
