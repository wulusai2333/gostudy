package main

/*
	socket编程:socket层是在应用层和传输层之间的一个抽象,相当于对应用封装屏蔽下层协议


	物理层 网线 传输10高低频信号
	数据链路层 广播数据,局域网中 网卡 Mac地址 mac包头14字节尾4字节
	网络层 网络管理员分配 IP地址 ip包 最大65535字节 头20字节 最大40字节
	传输层 UDP/TCP协议 端口号 UDP协议头8字节 TCP头20字节
	会话层 提供的服务是建立在维持会话，并且获得继续通信。通过校验点恢复通信
	表示层 为异种机器通信提供一种语言，以便能进行互相操作。因为不同计算机体系结构使用的数据表示法不同
	应用层 HTTP/FMTP

	在链路层，由以太网的物理特性决定了数据帧的长度为(46＋18)－(1500＋18)，其中的18是数据帧的头和尾，也就是说数据帧的内容最大为1500(不包括帧头和帧尾)
	在网络层，因为IP包的首部要占用20字节，所以这的MTU为1500－20＝1480
	在传输层，对于UDP包的首部要占用8字节，所以这的MTU为1480－8＝1472
	在应用层，你的Data最大长度为1472,UDP的数据最大为1472字节最好(避免分片重组)
	UDP 包的大小就应该是 1500 - IP头(20) - UDP头(8) = 1472(Bytes)
	TCP 包的大小就应该是 1500 - IP头(20) - TCP头(20) = 1460 (Bytes)
	鉴于Internet(非局域网)上的标准MTU值为576字节，所以建议在进行Internet的UDP编程时，最好将UDP的数据长度控制在548字节 (576-8-20)以内
*/
func main() {

}
