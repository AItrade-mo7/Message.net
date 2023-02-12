package disposeTask

import (
	"fmt"

	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTask"
	jsoniter "github.com/json-iterator/go"
)

// 发送完邮件之后要把发送记录存储到数据库中去
func SendSysEmail(opt any) {
	jsonByte := mJson.ToJson(opt)
	var EmailOpt mTask.SysEmail
	jsoniter.Unmarshal(jsonByte, &EmailOpt)

	fmt.Println("发送结束")
	mJson.Println(EmailOpt)
}

func SendCodeEmail(opt any) {
	fmt.Println("发送验证码")
}

func SendRegisterEmail(opt any) {
	fmt.Println("发送注册通知")
}

// 发送邮件的函数
// 用 gmail 发送
type EmailOpt struct {
	From     string
	To       []string
	Subject  string
	Template string
	SendData any
}

func FromGmail(opt EmailOpt) *mEmail.EmailInfo {
	emailObj := mEmail.New(mEmail.Opt{
		Account:     "meichangliang@gmail.com",
		Password:    "nmqlusfgaeyexxok",
		Port:        "587",
		Host:        "smtp.gmail.com",
		To:          opt.To,
		From:        opt.From,
		Subject:     opt.Subject,
		TemplateStr: opt.Template,
		SendData:    opt.SendData,
	})
	return emailObj
}

// 用 企业微信 发送

func FromWorkWeiXin(opt EmailOpt) *mEmail.EmailInfo {
	emailObj := mEmail.New(mEmail.Opt{
		Account:     "trade@mo7.cc",
		Password:    "DXir4WLb2aGaknLZ",
		Port:        "587",
		Host:        "smtp.exmail.qq.com",
		To:          opt.To,
		From:        opt.From,
		Subject:     opt.Subject,
		TemplateStr: opt.Template,
		SendData:    opt.SendData,
	})
	return emailObj
}

// 用 qq 发送
func FromQQ(opt EmailOpt) *mEmail.EmailInfo {
	emailObj := mEmail.New(mEmail.Opt{
		Account:     "meichangliang@qq.com",
		Password:    "fxdxnbyronppbfha",
		Port:        "587",
		Host:        "smtp.qq.com",
		To:          opt.To,
		From:        opt.From,
		Subject:     opt.Subject,
		TemplateStr: opt.Template,
		SendData:    opt.SendData,
	})
	return emailObj
}
