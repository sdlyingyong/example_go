package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	//defer conn.Close() 放这里会导致万一失败 没打开资源就关闭资源
	if err != nil {
		fmt.Println("listen UDP failed: err : ", err)
		return
	}
	defer conn.Close()
	//不需要建立连接,直接收发数据
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from UDP failed,err :", err)
			return
		}
		fmt.Println("收到消息: ", string(data[:n]))
		reply := "Reply: " + strings.ToUpper(string(data[:n]))
		//发送数据
		conn.WriteToUDP([]byte(reply), addr)
	}

}
