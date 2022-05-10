package main

import (
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
	_, err = c.Write([]byte("hi"))
	if err != nil {
		log.Fatal("write error:", err)
	}
	time.Sleep(100 * time.Millisecond)
}
