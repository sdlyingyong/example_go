package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	//替换default的设置
	//r := gin.Default()
	r := gin.New()

	//使用自定义的logger
	logger := InitLogger()

	//使用别人封装好的gin zap日志库
	r.Use(ZapLogger(logger), gin.Recovery())
	//r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	//注册接口并运行服务器
	r.GET("/hello", func(context *gin.Context) {
		logger.Debug(`context.String(200, "hello gin and zap") success`,
			zap.Any("context", context),
		)
		context.String(200, "hello gin and zap")
	})
	r.Run()
}

//初始化日志对象
func InitLogger() *zap.Logger {
	//准备配置参数
	enc := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	file, _ := os.Create("./zap.log")
	ws := zapcore.AddSync(file)
	enab := zapcore.DebugLevel

	//创建日志对象
	core := zapcore.NewCore(enc, ws, enab)
	logger := zap.New(core)
	//sugarLog := logger.Sugar()
	return logger
}

//中间件要实现handleFunc 类型
// HandlerFunc defines the handler used by gin middleware as return value.
//type HandlerFunc func(*Context)
func ZapLogger(zapLogger *zap.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		//记录参数收集
		start := time.Now()
		path := context.Request.URL.Path
		query := context.Request.URL.RawQuery
		context.Next()
		cost := time.Since(start)
		//日志信息写入
		zapLogger.Info(path,
			zap.String("writer", "ZapLogger"),
			zap.Int("status", context.Writer.Status()),
			zap.String("method", context.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", context.ClientIP()),
			zap.String("user-agent", context.Request.UserAgent()),
			zap.String("errors", context.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
