package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	humanize "github.com/dustin/go-humanize"
)

//下载工具
func main() {

	//参数处理
	if len(os.Args) < 2 {
		fmt.Println("请输入想要获取的链接")
		return
	}
	url := os.Args[1]
	fmt.Println("请求的链接是: ", url)
	startTime := time.Now()

	//网络请求,文件获取
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("fasthttp.Get(nil,url) failed, err :", err)
		return
	}
	defer resp.Body.Close()

	//文件存储
	//换用io.Copy,解决readAll大文件内存耗尽问题
	saveFile := fmt.Sprintf("tmp%s", path.Ext(url))
	file, err := os.Create(saveFile)
	if err != nil && err != io.EOF {
		fmt.Println("os.Create(saveFile) failed, err: ", err)
		return
	}
	defer file.Close()

	//写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil && err != io.EOF {
		fmt.Println("io.Copy(file,resp.Body) failed, err: ", err)
		return
	}

	//结果输出
	useTime := time.Since(startTime)
	fmt.Println("download success! ")
	f, err := os.Stat(saveFile)
	if err != nil {
		fmt.Println("os.Stat(saveFile) failed, err: ", err)
		return
	}
	fmt.Printf("use time : %s , size: %s \n", useTime, humanize.Bytes(uint64(f.Size())))
	os.Exit(0)
}

//许愿式编程
//直接就能协程把东西下载下来
//ty url.com/video.mp4

//0.0 支持命令行下载
//编译为ty.exe
//用法 ty.exe ../tmp.mp4

//0.1增加下载用时的统计
//0.2下载文件大小显示,kb,mb,gb
//0.3直观的运行时间
//0.4解决oom问题,换用io.copy

//测速
//测速链接(100MB)
//ty https://speedtest2.niutk.com:8080/download?size=100000000&r=0.5807917751032634

//todo
//限制占用内存 2G
//充分利用多核,开启协程多线程下载
//提高速度,也许用fast http

//build
//linux
//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ty main.go
