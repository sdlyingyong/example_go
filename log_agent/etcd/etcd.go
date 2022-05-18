package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`  //log file path
	Topic string `json:"topic"` //topic
}

func Init(addr string, timeout int) error {
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: time.Duration(timeout) * time.Second})
	if err != nil {
		fmt.Printf("conn etcd fail, err : %v/n", err)
		return err
	}
	return nil
}

//get conf from etcd
//json
//  c:/tmp/nginx.kig web_log
//	d://xxx/redis.log redis_log
func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get value from etcd fail, err :%v/n", err)
		return
	}
	for _, val := range resp.Kvs {
		err := json.Unmarshal(val.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value fail , err : %v/n", err)
		}
	}
	fmt.Println("get value from etcd success.")
	for index, val := range logEntryConf {
		fmt.Printf("index : %v value: %v \n", index, val)
	}

	return logEntryConf, err
}

func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)
	fmt.Println("开始监听配置热更新...")

	for wresp := range ch {
		for _, evt := range wresp.Events {
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("unmarshal fail, err : %v\n", err)
				}
			}
			fmt.Printf("new conf get by watch : %s\n", newConf)

			newConfCh <- newConf
		}
	}
}
