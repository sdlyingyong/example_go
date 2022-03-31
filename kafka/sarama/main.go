package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	//生产者
	kafkaProducer()
}

//kafka生产者案例
func kafkaProducer() {
	//kafka设置
	addr := []string{"127.0.0.1:9092"}
	config := sarama.NewConfig()
	// 发送完数据需要leader和follow都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 成功交付的消息将在success channel返回
	config.Producer.Return.Successes = true
	//连接kafka
	kfkProducer, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer closed ,err: ", err)
		return
	}
	fmt.Println("conn kafka success.")
	defer kfkProducer.Close()

	//构造发送的消息 topic => stringValue
	msg := &sarama.ProducerMessage{}
	msg.Topic = "example_go_api_log" //不能有空格
	msg.Value = sarama.StringEncoder("send msg success. pid:0 offset:3")

	//发送消息
	pid, offset, err := kfkProducer.SendMessage(msg)
	if err != nil {
		fmt.Println("kClient.SendMessage(msg) failed", err.Error())
		return
	}
	fmt.Printf("send msg success. pid:%v offset:%v \n", pid, offset)
}
