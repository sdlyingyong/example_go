package main

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	//服务器端口
	addr = "127.0.0.1:8080"
	//网卡序号
	netDeviceNo = 1
	//网卡包过滤条件
	filter = " "
)

//打印本机网卡收到的连接
func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("net.Listen failed, err: ", err)
		return
	}

	//等待客户端建立连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept failed, err: ", err)
			return
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	//接收数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read failed, err: ", err)
		return
	}
	fmt.Printf("read %d bytes, content is: %s\n", n, string(buf[:n]))
	showNetPacket()
}

func showNetPacket() {
	var (
		//device      string        = "eth0"
		snapshotLen int32         = 1024
		promiscuous bool          = false
		timeout     time.Duration = 30 * time.Second
		handle      *pcap.Handle
	)

	//获取网卡列表
	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Println("pcap.FindAllDevs() failed, err :", err)
		return
	}

	//接收网卡数据
	// Open device
	handle, err = pcap.OpenLive(devices[netDeviceNo-1].Name, snapshotLen, promiscuous, timeout)
	if err != nil {
		fmt.Println(" pcap.OpenLive(device filed, err: ", err)
		return
	}
	defer handle.Close()

	//过滤
	err = handle.SetBPFFilter(filter)
	if err != nil {
		return
	}

	//读取包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		//printIPPacket(packet)
		printUDPInfo(packet)
	}
}

//输出UDP报文的信息
func printUDPInfo(packet gopacket.Packet) (err error) {
	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer == nil {
		return
	}
	udp, ok := udpLayer.(*layers.UDP)
	if !ok {
		err = errors.New("not UDP format")
		return
	}
	// Check for errors
	if packet.ErrorLayer() != nil {
		fmt.Println("Error decoding some part of the packet:", err)
		return
	}
	//打印
	fmt.Printf("src_port: %s \n", udp.SrcPort)
	fmt.Printf("dst_port: %s \n", udp.DstPort)
	fmt.Printf("udp_length: %d \n", udp.Length)
	fmt.Printf("udp_checksum: %d \n", udp.Checksum)
	fmt.Println()
	return

}

//输出ip报文的信息
// IP layer variables:
// Version (Either 4 or 6)
// IHL (IP Header Length in 32-bit words)
// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
// Checksum, SrcIP, DstIP
func printIPPacket(packet gopacket.Packet) (err error) {
	// Let's see if the packet is IP (even though the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return
	}

	ip, ok := ipLayer.(*layers.IPv4)
	if !ok {
		err = errors.New("not ipv4 format")
		return
	}
	// Check for errors
	if packet.ErrorLayer() != nil {
		fmt.Println("Error decoding some part of the packet:", err)
		return
	}

	//打印
	fmt.Printf("%s to %s conn \n", ip.SrcIP.String(), ip.DstIP.String())
	fmt.Println("ip_version", ip.Version)
	fmt.Println("iph_length", ip.Length)
	fmt.Println("ttl", ip.TTL)
	fmt.Println("protocol", ip.Protocol)
	fmt.Println("iph_checksum", ip.Checksum)
	fmt.Println("src_ip", ip.SrcIP)
	fmt.Println("dst_ip", ip.DstIP)
	fmt.Println()

	return
}
