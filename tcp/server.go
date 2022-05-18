package main

import (
	"fmt"
	"net"
)

func main() {
	//创建服务器
	tcpServer()
}

//创建服务器
func tcpServer() {
	//监听端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	//接收客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}
		go handleConn(conn)
	}
}

//处理连接
func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("客户端断开连接")
				return
			}
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("Received:", string(buf[:n]))
	}
}
