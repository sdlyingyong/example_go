package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

func main() {
	r := gin.Default()
	//session开启
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//验证码接口
	r.GET("/captcha", captcha)
	//验证码检验
	r.GET("/verify", verify)

	//开启服务器
	r.Run(":8080")
	fmt.Println("开启服务器 8080")
}

//-------------------controller-------------------

func captcha(ctx *gin.Context) {
	//参数获取
	//获取sessionId
	sId := GetSessionId()
	//逻辑处理
	code, err := CaptchaHandle(sId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	//响应处理
	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{}, "msg": "OK"})
}

func verify(ctx *gin.Context) {
	//参数获取
	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "请输入登录验证码"})
		return
	}
	//获取sessionId
	sId := GetSessionId()
	//逻辑处理
	isPass := base64Captcha.VerifyCaptchaAndIsClear(sId, code, false) //验证错误不重新刷新
	//isPass := base64Captcha.VerifyCaptcha(sId, code)	//验证错误验证码重新生成
	if !isPass {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "验证码不正确 code:" + code})
		return
	}
	//响应处理
	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{}, "msg": "OK"})
}

//-------------------logic-------------------

//生成验证码
func CaptchaHandle(sessionId string) (code string, err error) {
	var configC = base64Captcha.ConfigCharacter{
		Height:             40,
		Width:              90,
		Mode:               base64Captcha.CaptchaModeNumber, //const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4,
	}
	_, capC := base64Captcha.GenerateCaptcha(sessionId, configC)
	//以base64编码
	code = base64Captcha.CaptchaWriteToBase64Encoding(capC)
	return
}

//-------------------utils-------------------
//需要按照框架获取sessionId
func GetSessionId() (sessionId string) {
	sessionId = "1001"
	return
}
