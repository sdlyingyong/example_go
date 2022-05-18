package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	//0.conn etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second})
	if err != nil {
		fmt.Printf("conn etcd fail, err : %v\n", err)
		return
	}
	fmt.Println("conn etcd success")
	defer cli.Close()
	//1.put key value to etcd
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "ty", "ttty")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd fail, err : %v/n", err)
	}
	fmt.Println("put key value to etcd success")
	//2.get from etcd
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "ty")
	cancel()
	if err != nil {
		fmt.Printf("get value from etcd fail, err : %v\n", err)
		return
	}
	for _, val := range resp.Kvs {
		fmt.Print("get value from etcd success ")
		fmt.Printf("%s:%s\n", val.Key, val.Value)
	}
	//3.watch one key->value change
	//执行对watch减值对的操作 会触发通知事件
	//etcdctl.exe --endpoints=http://127.0.0.1:2379 put ty "dsb2"
	//etcdctl.exe --endpoints=http://127.0.0.1:2379 del ty
	ch := cli.Watch(context.Background(), "ty")
	fmt.Println("wait watch event ...")
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type: %v key:%v value:%v \n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}
}
