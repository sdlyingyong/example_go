package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	//读取上传的excel文件
	fileName, err := ImportExcel()
	if err != nil {
		fmt.Println("ImportExcel() failed: ", err)
		return
	}
	//读取文件
	data, err := ReadExcel(fileName)
	if err != nil {
		return
	}

	fmt.Println("success! read rows :", len(data))
}

//导入excel
func ReadExcel(fileName string) (result []map[string]interface{}, err error) {
	//xlsx打开文件
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		return
	}
	//遍历每个表格
	result = make([]map[string]interface{}, 0)
	for _, sheet := range xlFile.Sheet {
		//每行,每列读取
		for index, row := range sheet.Rows {
			var gName, gStore string
			//过滤空行
			if index == 0 {
				continue
			}
			if len(row.Cells) < 2 {
				result = append(result, map[string]interface{}{"status": false, "data": bson.M{"group_name": "", "group_store": ""}, "error": "请填写分组名和门店编号"})
				continue
			}
			//获取分组名
			gName = strings.Trim(row.Cells[0].String(), " ")
			if gName == "" {
				result = append(result, map[string]interface{}{"status": false, "data": bson.M{"group_name": "", "group_store": ""}, "error": "请填写分组名"})
				continue
			}
			//获取门店编号
			gStore = strings.Trim(row.Cells[1].String(), "")
			if gStore == "" {
				result = append(result, map[string]interface{}{"status": false, "data": bson.M{"group_name": "", "group_store": ""}, "error": "请填写门店编号"})
				continue
			}
			//操作数据
			fmt.Println("分组名称: ", gName, " 店铺编号: ", gStore)
			result = append(result, map[string]interface{}{"status": false, "data": bson.M{"group_name": gName, "group_store": gStore}})
		}
	}
	return
}

//导入excel
func ImportExcel() (fileName string, err error) {
	//(框架)文件型参数存储在服务器目录
	//FileParamBeego()

	fileName = "./testDoc/excel_demo.xlsx" //测试文件
	return
}

//辅助函数

//获取上传的文件(beego版)
//github.com/astaxie/beego v1.12.3
func FileParamBeego() {
	////参数校验
	//_, head, err := this.GetFile("excel")
	//if err != nil {
	//	this.Result(0, nil, err.Error())
	//	return
	//}
	//if head.Filename[len(head.Filename)-4:] != "xlsx" {
	//	this.Result(0, nil, "请上传xlsx格式的数据文档")
	//	return
	//}
	//saveUrl := "/risks/excel/" + time.Now().Format("200601") + "/"
	//err = utils.IsFile(saveUrl)
	//if err != nil {
	//	this.Result(0, nil, err.Error())
	//	return
	//}
	//bsonId := bson.NewObjectId().Hex()
	//fileName := saveUrl + bsonId + ".xlsx"
	//err = this.SaveToFile("excel", fileName)
	//if err != nil {
	//	this.Result(0, nil, "保存excel文件失败")
	//	return
	//}
}

//判断文件夹是否存在
func IsFile(name string) (err error) {
	_, err = os.Stat(name)
	if os.IsNotExist(err) {
		err = os.MkdirAll(name, os.ModePerm)
	}
	return
}
