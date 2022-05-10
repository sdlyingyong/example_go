package main

import (
	"log"
	"net"
	"os"
)

const SockAddr = "/tmp/echo.sock"

//显示通过socket连接的client
func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	buf := make([]byte, 1024)
	n, err := c.Read(buf[:])
	if err != nil {
		return
	}
	println("Client send msg:", string(buf[0:n]))
	//关闭本次连接
	c.Close()
}

func main() {
	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	//监听socket
	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	//监听消息
	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(conn)
	}
}
