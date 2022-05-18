package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type LogData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer
	err         error
	logDataChan chan *LogData
)

func Init(addr []string, maxSize int) error {
	//ready config
	config := sarama.NewConfig()

	//ack type = all
	//need leader and followers return ack
	config.Producer.RequiredAcks = sarama.WaitForAll
	//new choose one partition
	//type is foreach request : sarama.NewRandomPartitioner
	//type is key,use hash WithCustomHashFunction
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// success msg return in success channel
	config.Producer.Return.Successes = true

	//conn kafka
	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer closed ,err: ", err)
		return err
	}
	//fmt.Println("conn kafka success.")

	//create log send chan
	logDataChan = make(chan *LogData, maxSize)
	//start another goroutine : get data from chan and send to kafka
	go ListenAndSendToKafka(logDataChan)
	return nil
}

//get data outside and  send to inside chan
func SendToChan(topic, data string) {
	msg := &LogData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

//send to kafka
func ListenAndSendToKafka(logDataChan chan *LogData) {
	for {
		select {
		case ld := <-logDataChan:
			//write msg for kafka
			msg := &sarama.ProducerMessage{}
			// get data from chan
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			//send msg
			_, _, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err :", err)
				return
			}
			fmt.Println("send to kafka success. log data :", ld)
		default:
			time.Sleep(time.Millisecond * 50)
		}

	}
}
