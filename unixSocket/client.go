package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	//socket连接
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	//消息写入
	msg := "hi"
	_, err = c.Write([]byte(msg))
	if err != nil {
		log.Fatal("write error:", err)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("client send msg :", msg)
}
