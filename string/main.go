package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

//demo 为案例
//exec 为执行函数
func main() {

	//强密码检查
	demoCheckStrongPwd()

	//极光推送
	demoJPush()

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

func demoJPush() {
	pushTags := make([]string, 0)
	pushTags = append(pushTags, "id-xxxx") //设备号
	ret, err := PushExec(pushTags, nil, "反馈处理", "您的反馈已经处理，请查看！", 2, "")
	if err != nil {
		fmt.Println("demoJPush failed, err: ", err)
		return
	}
	fmt.Println("demoJPush success, ret: ", ret)
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

//推送消息
//@param tags 推送用户登录的token（用户登录被挤掉）没有可为nil 不能和alias同时为空
//@param alias 推送用户所在组织的ID 没有可为nil 不能和tags同时为空
//@param $type 消息类型 0=IOS APP版本更新 1=安卓版本更新 2=反馈回复 其他待定
//@param msgId 消息ID
type M map[string]interface{} //自定义map
func PushExec(tags []string, alias []string, title, alert string, _type uint8, msgId string) (ret interface{}, err error) {

	//准备极光的接口认证参数
	req := httplib.Post("https://api.jpush.cn/v3/push")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("Content-Type", "application/json")
	appkey := ""                                                 //极光推送给你提供的appkey
	secret := ""                                                 //极光推送给你提供的的secret
	isProduct := beego.AppConfig.DefaultBool("isProduct", false) //是否产品模式
	req.SetBasicAuth(appkey, secret)

	//推送的对象设置
	var param = M{}
	param["platform"] = "all"
	audience := M{}
	if tags != nil {
		audience["tag"] = tags
	}
	if alias != nil {
		audience["alias"] = alias
	}
	param["audience"] = audience
	var content = make(M)
	if _type == 0 {
		content["ios"] = M{
			"alert":             M{"title": title, "body": alert},
			"sound":             "default",
			"badge":             "+1",
			"extras":            M{"msgId": msgId, "type": _type},
			"content-available": true,
		}
	} else if _type == 1 {
		content["android"] = M{
			"alert":      alert,
			"title":      title,
			"builder_id": 1,
			"extras":     M{"msgId": msgId, "type": _type},
		}
	} else {
		content["ios"] = M{
			"alert":             M{"title": title, "body": alert},
			"sound":             "default",
			"badge":             "+1",
			"extras":            M{"msgId": msgId, "type": _type},
			"content-available": true,
		}
		content["android"] = M{
			"alert":      alert,
			"title":      title,
			"builder_id": 1,
			"extras":     M{"msgId": msgId, "type": _type},
		}
	}
	param["notification"] = content
	param["message"] = M{
		"msg_content":  alert,
		"content_type": "text",
		"title":        title,
		"extras":       M{"msgId": msgId, "type": _type}}
	param["options"] = M{"apns_production": isProduct, "time_to_live": 60}

	//发送请求并返回结果
	a, _ := json.Marshal(param)
	req.Body(string(a))
	data, err := req.Bytes()
	if err != nil {
		return
	}
	ret = struct {
		Sendno string
		Msg_id string
	}{}
	json.Unmarshal(data, &ret)
	return
}
