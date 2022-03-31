package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/juju/ratelimit"
)

//接口限制两秒访问一次
func main() {
	r := gin.Default()
	//创建令牌桶
	bucket := ratelimit.NewBucket(2*time.Second, 1) //2秒放一个令牌
	r.GET("/ping", func(c *gin.Context) {
		//每次从令牌桶里拿,没有就返回限速
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit...")
			return
		}
		//有就放行
		c.String(http.StatusOK, "hello")
	})
	//开启服务器
	r.Run(":8080")
}
