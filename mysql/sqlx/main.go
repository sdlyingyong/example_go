package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //注入方式使用mysql
	"github.com/jmoiron/sqlx"
)

var (
	//全局连接
	client = initDb()
)

//定义结构体 对应数据库中的table
//1.需要首字母大写
//2.需要结构体db:"字段"
type User struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

//初始化mysql连接
func initDb() (db *sqlx.DB) {
	dsn := "root:root@tcp(127.0.0.1:3306)/sqlx?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//必须连接 出问题就直接panic
	db = sqlx.MustConnect("mysql", dsn)

	//最大连接数
	db.SetMaxOpenConns(20)
	//最大处理数
	db.SetMaxIdleConns(10)
	fmt.Println("conn mysql success")
	return
}

func main() {

	//插入数据
	namedExec()
	//批量插入
	batchInsert()

	//修改
	updateRow()

	//查询数据库
	sqlStr := "SELECT id,name ,age FROM users WHERE id = ?"
	var u1 User
	err := client.Get(&u1, sqlStr, 1) //传指针&u1

	if err == sql.ErrNoRows {
		fmt.Println("client.Get(&u1, sqlStr, 1) success,but ret null")
	} else if err != nil {
		fmt.Println("client.Get(&u1, sqlStr, 1) failed,err :", err)
	} else {
		fmt.Println("client.Get success, u1 is", u1)
	}

	//查询多条
	queryMultiRow()

	//查询(结构体传参)
	namedQuery()
}

//查询多条
func queryMultiRow() {
	sqlStr := "select id ,name, age from users where id > ?"
	var ulist []User
	err := client.Select(&ulist, sqlStr, 0)
	if err != nil {
		fmt.Println("query fail,err :", err)
		return
	}
	fmt.Println("queryMultiRow success, ret :", ulist)
}

//修改
func updateRow() {
	sqlStr := "update users set age = ? where id = ?"
	//delStr := "delete from user where id = ? "
	//insertStr := "insert into user (name,age) values(?)"
	ret, err := client.Exec(sqlStr, 23, 2)
	if err != nil {
		fmt.Println("update fail,err :", err)
		return
	}
	nums, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("get rows affected fail, err :", err)
		return
	}
	fmt.Println("updateRow success,ret :", nums)
}

//结构体方式插入数据
func namedExec() {
	insertSrt := `insert into users (name,age) values (:name,:age)`
	insertData := map[string]interface{}{
		"name": "七米",
		"age":  32,
	}
	ret, err := client.NamedExec(insertSrt, insertData)
	if err != nil {
		fmt.Println("insert fail, err :", err)
		return
	}
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("get rows affected fail, err :", err)
		return
	}
	fmt.Println("namedExec success, insert row num is ", num)
}

//结构参数查询
func namedQuery() {
	//用map格式查询条件
	sqlStr := "SELECT * FROM users WHERE name=:name"
	queryData := map[string]interface{}{
		"name": "七米",
	}
	rows, err := client.NamedQuery(sqlStr, queryData)
	if err != nil {
		fmt.Println("namedQuery fail, err :", err)
		return
	}
	defer rows.Close()
	//用结构体查询条件
	u := User{Name: "七米"}
	rows2, err := client.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Println("namedQuery fail, err :", err)
		return
	}
	defer rows2.Close()
	//遍历取出
	for rows.Next() {
		var getU User
		err := rows.StructScan(&getU)
		if err != nil {
			fmt.Println("StructScan fail, err : ", err)
			return
		}
		//fmt.Printf("user: %#v \n", getU)
	}

	fmt.Println("namedQuery success")

}

//准备批量数据
func batchInsert() {
	u1 := User{0, 28, "七米"}
	u2 := User{0, 18, "qimi"}
	u3 := User{0, 38, "小王子"}
	users := []interface{}{u1, u2, u3}
	err := batchInsertExec(users)
	if err != nil {
		fmt.Println("batchInsertExec(users) failed, err", err.Error())
		return
	}
	fmt.Println("batchInsert success")
}

//实现driver方法的value,用于sqlx.In
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

//批量插入
func batchInsertExec(users []interface{}) (err error) {
	//需要我们的结构体user实现driver.Valuer
	query, args, _ := sqlx.In(
		"INSERT INTO users (name,age) VALUES (?),(?),(?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	_, err = client.Exec(query, args...)
	if err != nil {
		fmt.Println("client.Exec(query, args...) failed, err: ", err)
		return
	}
	return
}
