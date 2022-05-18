package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//router
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.String(200, "Hello , TY")
	})
	//curl接口注册处理请求
	//请求地址中的参数取出
	r.GET("/task/:taks_id", func(context *gin.Context) {
		task_id := context.Query("task_id")
		context.String(http.StatusOK, "get task id : "+task_id+" detail :orm select and return result")
	})
	r.DELETE("/task/:task_id", func(context *gin.Context) {
		task_id := context.Query("task_id")
		context.String(http.StatusOK, "delete task id is :"+task_id+" orm delete it and return 200")
	})
	r.POST("/task", func(context *gin.Context) {
		task_name := context.PostForm("task_name")
		task_type := context.PostForm("task_type")
		test_id := context.PostForm("test_id")
		tester_user_id := context.PostForm("tester_user_id")
		reviewer_user_id := context.PostForm("reviewer_user_id")
		//some orm insert
		context.JSON(http.StatusOK, gin.H{
			"task_name":        task_name,
			"task_type":        task_type,
			"test_id":          test_id,
			"tester_user_id":   tester_user_id,
			"reviewer_user_id": reviewer_user_id,
		})
	})
	//返回json格式响应
	r.PUT("/task/:task_id", func(context *gin.Context) {
		task_id := context.Param("task_id")
		task_name := context.PostForm("task_name")
		task_type := context.DefaultPostForm("task_type", "simple")
		test_id := context.PostForm("test_id")
		tester_user_id := context.PostForm("tester_user_id")
		reviewer_user_ids := context.PostFormArray("reviewer_user_ids")
		//some orm update
		context.JSON(http.StatusOK, gin.H{
			"task_id":           task_id,
			"task_name":         task_name,
			"task_type":         task_type,
			"test_id":           test_id,
			"tester_user_id":    tester_user_id,
			"reviewer_user_ids": reviewer_user_ids,
		})
	})
	//接口使用json数据
	r.PUT("/json", func(context *gin.Context) {
		jsonBytes, _ := context.GetRawData()
		result := make(map[string]interface{})
		json.Unmarshal(jsonBytes, &result)
		context.JSON(http.StatusOK, result)
	})
	//重定向到指定页面
	r.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/welcome")
	})
	r.GET("/welcome", func(context *gin.Context) {
		context.String(200, "hello now time is %s", time.Now())
	})
	//分组路由
	//用于区分接口v1 /v2 或者分区加权限检查
	noAuth := r.Group("/noAuth")
	{
		noAuth.POST("/login", func(context *gin.Context) {
			context.String(http.StatusOK, "this login not need auth")
		})
		noAuth.GET("/version", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"version": "v1",
			})
		})
	}
	//上传单个文件
	r.POST("upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		fmt.Println("Filename is : ", file.Filename)
		dst := fmt.Sprintf("D:/tmp/%s", file.Filename)
		context.SaveUploadedFile(file, dst)
		context.JSON(http.StatusOK, gin.H{
			"msg": "file name :" + file.Filename,
		})
	})
	//上传多个文件
	r.POST("/uploadFiles", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			//抛出错误,gin.Default()->engine.Use(Recovery()) 会捕获并转为500
			panic(err)
		}
		files := form.File["file"] //param name
		result := make([]string, 0)
		for _, file := range files {
			result = append(result, file.Filename)
			dst := fmt.Sprintf("D:/tmp/%s", file.Filename)
			context.SaveUploadedFile(file, dst)
		}
		context.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})
	//每个请求都使用中间件处理
	//api处理创建新的goroutine时,要通过 Context.Copy() 制作新的上下文来使用
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//登录后才能使用
	auth := r.Group("/auth", gin.BasicAuth(gin.Accounts{"root": "s123"}))
	//需要登陆使用的接口全部使用安全检查中间件
	auth.Use(TokenCheck())
	{
		auth.GET("/my_info", func(context *gin.Context) {

			cardList := make([]string, 10)
			for i := 0; i < 10; i++ {
				cardList = append(cardList, "card No. "+strconv.Itoa(i))
			}

			curr_user, _ := context.Get("user_info")

			context.JSON(http.StatusOK, gin.H{
				"name":      "zhansan",
				"cards":     cardList,
				"user_info": curr_user,
			})
		})
	}
	//处理不正确的访问端口,提示错误
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"msg":  "check you api endpoint",
			"code": "404",
			"data": []string{},
		})
	})
	//开启服务器
	r.Run(":999")

}

//身份检查
func TokenCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		//请求处理前 before action
		sTime := time.Now()
		context.Set("user_info", gin.H{"name": "zhangsan", "email": "zs@gmail.com"})
		//执行请求
		context.Next()
		//请求后 after action
		fmt.Println("请求耗时: ", time.Since(sTime))
	}
}
