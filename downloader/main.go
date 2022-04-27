package main

import (
	"fmt"
	"io/ioutil"
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
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll(resp.Body) failed, err :", err)
		return
	}

	//根据url获取文件格式
	saveFile := fmt.Sprintf("tmp%s", path.Ext(url))
	err = ioutil.WriteFile(saveFile, data, 0644)
	if err != nil {
		fmt.Println(`ioutil.WriteFile("tmp", data, 0644)failed, err : `, err)
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
	fmt.Printf("use time : %s , size: %s", useTime, humanize.Bytes(uint64(f.Size())))
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

//测速
//测速链接(100MB)
//ty https://speedtest2.niutk.com:8080/download?size=100000000&r=0.5807917751032634

//todo
//限制占用内存 2G
// 充分利用多核,开启协程多线程下载
//提高速度,也许用fast http
