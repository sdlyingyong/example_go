package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//分组门店数据
type GroupStore struct {
	GroupName string
	Store     string
	Code      string
}

func main() {
	gsList := make([]*GroupStore, 0, 2) //长度0,容量2
	gsList = append(gsList, &GroupStore{
		GroupName: "DVR-0315-43",
		Store:     "816286215200777",
		Code:      "0001",
	})
	gsList = append(gsList, &GroupStore{
		GroupName: "DVR-0315-44",
		Store:     "816286215200778",
		Code:      "0002",
	})

	//写入excel
	filePath, err := WriteExcel(gsList)
	if err != nil {
		fmt.Println("WriteExcel() failed: ", err)
		return
	}
	fmt.Println("生成excel成功 filePath :", filePath)

	////导出下载文件
	//rs, err := DownloadExcel()
	//if err != nil {
	//	fmt.Println("DownloadExcel() failed: ", err)
	//	return
	//}
}

//导出excel
func WriteExcel(list []*GroupStore) (filePath string, err error) {
	//导出文件创建
	f := excelize.NewFile()
	sheet := "Sheet1"

	//excel样式设置
	f.SetRowHeight(sheet, 1, 30)
	f.SetColWidth(sheet, "A", "J", 20)
	f.SetColWidth(sheet, "E", "H", 15)
	headStyle, _ := f.NewStyle(`{"font":{"bold":true,"size":12,"family":"arial"},"alignment":{"vertical":"center"}}`)
	f.SetCellStyle(sheet, "A1", "J1", headStyle)

	//标题设置
	f.SetCellValue(sheet, "A1", "分组名称")
	f.SetCellValue(sheet, "B1", "门店编号")
	f.SetCellValue(sheet, "C1", "随机数")

	//写入每行数据
	line := 2
	for _, val := range list {
		f.SetRowHeight(sheet, line, 20)
		f.SetCellValue(sheet, fmt.Sprintf("A%d", line), val.GroupName)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", line), val.Store)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", line), val.Code)
		line++
	}

	//文件存储
	nowDate := time.Now().Format("200601")
	saveUrl := fmt.Sprintf("./%s/", nowDate)
	if err = IsFile(saveUrl); err != nil {
		err = errors.New("创建导入目录失败")
		return
	}
	fileName := "分组门店导出" + Time2Str(uint32(time.Now().Unix()), "20060102") + ".xlsx"
	if err = f.SaveAs(saveUrl + fileName); err != nil {
		err = errors.New("写入excel文件失败")
		return
	}
	filePath = fmt.Sprintf("excel/%s/%s", nowDate, fileName)
	return
}

//辅助函数

//判断文件夹是否存在
func IsFile(name string) (err error) {
	_, err = os.Stat(name)
	if os.IsNotExist(err) {
		err = os.MkdirAll(name, os.ModePerm)
	}
	return
}

//时间转字符串
func Time2Str(datetime uint32, format string) string {
	tm := time.Unix(int64(datetime), 0)
	return tm.Format(format)
}
