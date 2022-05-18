package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func (st *Student) run() *Student {
	fmt.Println(st.Name, "is running")
	return st
}

func (st *Student) wang() *Student {
	fmt.Println(st.Name, "is wang")
	return st
}

func showListCall() {
	//链式操作
	luminghui := Student{
		Name:    "卢明辉",
		Age:     1000,
		Married: false,
	}
	luminghui.run().wang().run().wang()
}

func main() {

	//use es client
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}

	p1 := Student{Name: "rion", Age: 22, Married: false}
	put1, err := client.Index().
		Index("user").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Indexed student : ", put1.Id, put1.Index, put1.Type)

}
