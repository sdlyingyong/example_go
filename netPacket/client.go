package main

import (
	"fmt"
	"net"
)

func main() {
	//连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Dial failed, err: ", err)
		return
	}
	//发送数据
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		fmt.Println("conn.Write failed, err: ", err)
		return
	}
	//接收数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read failed, err: ", err)
		return
	}
	fmt.Printf("read %d bytes, content is: %s\n", n, string(buf[:n]))
}
