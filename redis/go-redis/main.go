package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var (
	rdb *redis.Client
)

//初始化redis连接
func InitDb() (err error) {
	//数据库配置
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		errors.New("conn redis fail")
		fmt.Println("conn redis fail, err :", err)
		return
	}

	//另一个数据库使用
	var TwoRdb *redis.Client
	TwoRdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       2, //DB 2
	})
	_, err = TwoRdb.Ping().Result()
	if err != nil {
		fmt.Println("TwoRdb.Ping().Result(), err :", err)
		return
	}
	fmt.Println("InitDb success")
	return
}

func main() {
	//初始化redis连接
	InitDb()

	//获取从高到低的排名
	key := "rank"
	items := []redis.Z{
		redis.Z{Score: 90, Member: "PHP"},
		redis.Z{Score: 80, Member: "JAVA"},
		redis.Z{Score: 70, Member: "Go"},
	}
	num, err := rdb.ZAdd(key, items...).Result()
	if err != nil {
		fmt.Printf("zset err : %v \n", err)
		return
	}
	fmt.Println("rdb.ZAdd success, zset num :", num)

	//ZREVRANGEBYSCORE 返回有序集合中指定分数区间内的成员，分数由高到低排序。
	retS, err := rdb.ZRevRangeByScore(key, redis.ZRangeBy{"60", "100", 0, 10}).Result()
	if err != nil {
		fmt.Println("get by score err: ", err)
		return
	}
	fmt.Println("获取从高到低的排名 success, 按照热度排序的列表是: ", retS)

	//pipeline多个命令一起执行
	pipeline()
	pipelined()

	//多个命令(事务)
	txPipeline()
	txPipelined()

	//用来确保键的值不会被其他人修改,操作赋值
	//如果被人修改删除替换,就会收到一个错误
	watch()
	//修改值的同时被别人修改的案例
	watchExample()
}

//pipeline 多个命令一起执行
func pipeline() {
	//不互相依赖的操作可以用pipeline 一起发送到redis,一次执行
	pipe := rdb.Pipeline()

	incr := pipe.Incr("pipeline_counter")      //原值+1
	pipe.Expire("pipeline_counter", time.Hour) //设置过期时间

	_, err := pipe.Exec()
	if err != nil {
		fmt.Println("pipe.Exec() failed, err : ", err)
		return
	}
	fmt.Println("showPipeline success, ret ", incr.Val())
}

//pipelined 多个命令一起执行
func pipelined() {
	var incr *redis.IntCmd
	_, err := rdb.Pipelined(func(pipeliner redis.Pipeliner) error {
		//执行原子+1操作
		incr = pipeliner.Incr("pipelined_counter")
		//设置一个key的过期的秒数
		pipeliner.Expire("pipelined_counter", time.Hour)
		return nil
	})
	if err != nil {
		fmt.Println("rdb.Pipelined() failed, err : ", err)
		return
	}
	fmt.Println("showPipelined success, ret ", incr.Val())
}

//多个指令(事务)
func txPipeline() {
	pipe := rdb.TxPipeline()

	incr := pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	if err != nil {
		fmt.Println("pipe.Exec() failed, err : ", err)
		return
	}
	fmt.Println("showTxPipeline success, ret: ", incr.Val())
}

//多个指令(事务)
func txPipelined() {
	var incr *redis.IntCmd
	_, err := rdb.TxPipelined(func(pipeliner redis.Pipeliner) error {
		incr = pipeliner.Incr("tx_pipelined_counter")
		pipeliner.Expire("tx_pipelined_counter", time.Hour)
		return nil
	})
	if err != nil {
		fmt.Println("pipe.Exec() failed, err : ", err)
		return
	}
	fmt.Println("showTxPipeline success, ret: ", incr.Val())
}

//用来确保键的值不会被其他人修改,操作赋值
//如果被人修改删除替换,就会收到一个错误
func watch() (err error) {
	key := "watch_count"
	err = rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		tx.Pipelined(func(pipeliner redis.Pipeliner) error {
			pipeliner.Set(key, n+1, time.Minute*5)
			return nil
		})
		return err
	}, key)
	return
}

//修改值的同时被别人修改的案例
func watchExample() {
	const routineCount = 100

	// Increment 使用GET和SET命令以事务方式递增Key的值
	increment := func(key string) error {
		// 事务函数
		txf := func(tx *redis.Tx) error {
			// 获得key的当前值或零值
			n, err := tx.Get(key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际的操作代码（乐观锁定中的本地操作）
			n++

			// 操作仅在 Watch 的 Key 没发生变化的情况下提交
			_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
				// pipe handles the error case
				pipe.Set(key, n, 0)
				return nil
			})
			return err
		}

		//最多重试 maxRetries 次
		for retries := routineCount; retries > 0; retries-- {
			err := rdb.Watch(txf, key)
			if err != redis.TxFailedErr {
				return err
			}
			//优化的锁丢失
		}
		return errors.New("increment reached maximum number of retries")
	}

	//模拟 routineCount 个并发同时去修改 counter3 的值
	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()

			//并发安全地将redis值+1
			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := rdb.Get("counter3").Int()
	if err != nil {
		fmt.Println(`rdb.Get("counter3").Int() failed`, err)
		return
	}
	fmt.Println("watchExample success, ended with", n)
}
