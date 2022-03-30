package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	// 每秒放行的次数
	rl := ratelimit.New(20) //一秒放行20个 平均每个50ms

	//循环十次 打印每次的间隔时间
	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(time.Now())
		fmt.Println(i, now.Sub(prev))
		prev = now
	}

	//每个间隔执行一次操作
	for {
		rl.Take()
		fmt.Println("exec handle time ", time.Now())
	}
}
