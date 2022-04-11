package main

import (
	"fmt"
	"net"
)

func main() {
	//创建tcp连接
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Dial failed, err :", err)
		return
	}
	defer conn.Close()
	//发送数据
	_, err = fmt.Fprintf(conn, "Hello world \n")
	if err != nil {
		fmt.Println(`fmt.Fprintf(conn, "Hello world \n") failed, err :`, err)
		return
	}
	fmt.Println(`fmt.Fprintf(conn, "Hello world \n") success`)
}
