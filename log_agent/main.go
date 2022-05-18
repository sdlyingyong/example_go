package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"gopkg.in/ini.v1"
	"log_agent_ty/config"
	"log_agent_ty/etcd"
	"log_agent_ty/kafka"
	taillog "log_agent_ty/tail"
	"log_agent_ty/utils"
	"sync"
)

var (
	tailObj *tail.Tail
	LogChan chan string
	cfg     = new(config.AppConf)
)

func main() {
	//0.load conf
	//0.加载配置文件
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Println("load ini fail, err : ", err)
		return
	}

	//1.start kafka conn
	//并开启一个监听的协程:接收logChan 传来的内容
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("Init kafka conn fail, err :%v\n", err)
		return
	}
	fmt.Println("Init kafka conn success.")

	//2.start etcd conn
	err = etcd.Init(cfg.EtcdConf.Address, cfg.EtcdConf.Timeout)
	if err != nil {
		fmt.Printf("Init etcd conn fail, err : %v/n", err)
		return
	}
	fmt.Println("Init etcd conn success.")

	//2.1 访问etcd拿到监听日志的路径
	ip := utils.GetOutboundIP()
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key,ip)
	fmt.Println("get etcd conf key is :", etcdConfKey)

	//为了实现每个logagent拿到自己的配置,根据ip区分
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("get conf from etcd fail, err : %v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success  \n")

	//read file and send to kafka
	taillog.Init(logEntryConf)

	//watch etcd conf update
	newConfCh := taillog.NewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	//2.2 set a watch from log conf changes,and return .to bu hot update
	go etcd.WatchConf(cfg.EtcdConf.Key, newConfCh)
	wg.Wait()

}
