package main

import (
	"example_go/config/viper/conf"
	"fmt"
)

func main() {

	//配置加载器
	if err := conf.Init(); err != nil {
		fmt.Println("Init settings failed, err :", err)
		return
	}

	//配置值获取
	mysqlConf := conf.Conf.MysqlConfig
	fmt.Printf("mysql配置信息为: host:%s port:%s user:%s password:%s db:%s \n",
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Dbname,
	)
}
