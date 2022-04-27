package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
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
	startTime := time.Now().UnixNano()

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
	err = ioutil.WriteFile(fmt.Sprintf("tmp%s", path.Ext(url)), data, 0644)
	if err != nil {
		fmt.Println(`ioutil.WriteFile("tmp", data, 0644)failed, err : `, err)
		return
	}
	//结果输出
	useTime := float64(time.Now().UnixNano()-startTime) / (1000.000 * 1000.00)
	fmt.Println("download success! use time："+strconv.FormatFloat(useTime, 'f', -1, 64), "ms")
}

//编译为ty.exe
//用法 ty.exe ../tmp.mp4

//许愿式编程
//直接就能协程把东西下载下来
//ty baidu.com/file.mp4

//0.1增加下载用时的统计

//todo
//限制占用内存 2G

// 充分利用多核,开启协程多线程下载
