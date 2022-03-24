package main

import (
	"errors"
	"fmt"
	"regexp"
)

//demo 为案例
//exec 为执行函数
func main() {

	//强密码检查
	demoCheckStrongPwd()

}

//强密码案例
func demoCheckStrongPwd() {
	invalidStr := "12345678"
	if err := CheckStrongPwdExec(invalidStr); err == nil {
		fmt.Println("CheckStrongPwdExec(validStr) failed")
		return
	}
	validStr := "123456Qw!"
	if err := CheckStrongPwdExec(validStr); err != nil {
		fmt.Println("CheckStrongPwd(validStr) failed, err: ", err)
		return
	}
	fmt.Println("demoCheckStrongPwd success")
}

//检测强密码
//是否符合规则:至少八位,包含大小写字母,数字,符号
func CheckStrongPwdExec(pwd string) error {
	//不能少于8位
	if len(pwd) < 8 {
		return errors.New("密码不能少于8位")
	}
	//必须有大写字母
	if reg := regexp.MustCompile(`[A-Z]+`); !reg.MatchString(pwd) {
		return errors.New("密码必须包含大写字母")
	}
	//小写字母
	if reg := regexp.MustCompile(`[a-z]+`); !reg.MatchString(pwd) {
		return errors.New("密码必须包含小写字母")
	}
	//数字
	if reg := regexp.MustCompile(`[0-9]+`); !reg.MatchString(pwd) {
		return errors.New("密码必须包含数字")
	}
	//符号
	if reg := regexp.MustCompile(`[~!@#$%^&*()_\-+=<>?:"{}|,.\/;'\\[\]·~！@#￥%……&*（）——\-+={}|《》？：“”【】、；‘'，。、]`); !reg.MatchString(pwd) {
		return errors.New("密码必须包含特殊符号")
	}
	return nil
}
