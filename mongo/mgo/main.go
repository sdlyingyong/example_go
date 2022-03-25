package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

//全局数据库连接对象
var (
	MongoDb *mgo.Database
)

//mongo (表)结构体定义
type Devices struct {
	// mongo字段定义 bson:""
	Device_id   uint64 `bson:"device_id" json:"device_id" `
	Device_name string `bson:"device_name" json:"device_name" `
	Created_at  int64  `bson:"created_at" json:"created_at" `
}

func main() {
	//初始化连接
	MongoDb, err := InitMongo()
	if err != nil {
		fmt.Println("InitMongo() failed , err: ", err)
		return
	}
	fmt.Println("InitMongo success")

	//集合选择器
	devModel := MongoDb.C("devices")
	//集合插入操作
	newDev := &Devices{
		Device_id:   1,
		Device_name: "example_go的设备",
		Created_at:  time.Now().Unix(),
	}
	if err = devModel.Insert(newDev); err != nil {
		fmt.Println("InitMongo() failed , err: ", err)
		return
	}
	//操作结果
	fmt.Println("devModel.Insert(date) success")
}

//初始化连接
func InitMongo() (MongoDb *mgo.Database, err error) {
	//配置信息
	host := "127.0.0.1"
	dbName := "example_go"
	user := ""
	pwd := ""
	info := &mgo.DialInfo{
		// []string{host}
		Addrs:    []string{host},
		Database: dbName,
		Timeout:  60 * time.Second,
		Username: user,
		Password: pwd,
	}
	//数据库连接
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		//终止并退出
		log.Fatal("mangodb1:", err)
	}
	session.SetMode(mgo.Eventual, true)
	//返回连接对象
	MongoDb = session.DB("")
	return
}
