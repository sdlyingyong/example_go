package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {

	//conn etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second})

	if err != nil {
		fmt.Printf("conn etcd fail, err : %v\n", err)
		return
	}
	fmt.Println("conn etcd success")
	//close conn later
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//put key value to etcd
	key := "/loagent/192.168.2.9/collection_config"
	//key := "/loagent/collection_config"
	//value := `[{"path":"D:/tianyou/dev/log_agent_ty/logs/redis.txt","topic":"redis_log"}]`
	//value := `[{"path":"D:/tianyou/dev/log_agent_ty/logs/web.txt","topic":"web_log"},{"path":"D:/tianyou/dev/log_agent_ty/logs/redis.txt ","topic":"redis_log"}]`
	value := `[{"path":"D:/tianyou/dev/log_agent_ty/logs/web.txt","topic":"web_log"},{"path":"D:/tianyou/dev/log_agent_ty/logs/redis.txt","topic":"redis_log"},{"path":"D:/tianyou/dev/log_agent_ty/logs/mysql.txt","topic":"mysql_log"}]`
	_, err = cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd fail, err : %v/n", err)
	}
	fmt.Printf("put to etcd success, key : %v /n",key)

}
