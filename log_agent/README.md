日志监控项目

  通过kafka发送到ELK

1.需要运行服务

	kafka etcd
  
2.安装go依赖 

	go1.16
	go mod tidy
  
3.给etcd增加配置文件,

	去etcd/put/ 运行go run main.go
  
4.开始监控日志文件	

	/目录 go run main.go
