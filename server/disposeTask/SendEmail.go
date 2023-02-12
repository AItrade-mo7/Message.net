package disposeTask

import (
	"fmt"

	"Message.net/server/tmpl"
	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTask"
	jsoniter "github.com/json-iterator/go"
)

// 发送完邮件之后要把发送记录存储到数据库中去
func SendSysEmail(opt any) {
	jsonByte := mJson.ToJson(opt)
	var info mTask.SendEmail
	jsoniter.Unmarshal(jsonByte, &info)

	emailObj := BuildEmail(EmailOpt{
		From:     info.From,
		To:       info.To,
		Subject:  info.Subject,
		TmplName: info.TmplName,
		SendData: info.SendData,
	})

	err := emailObj.Send()
	fmt.Println("邮件发送完成", err)
}

// ======== 发件池子   ==============
var EmailAccountList = []mEmail.ServeType{
	mEmail.Gmail("mo7trade1@gmail.com", "bhmfbovjxnkmcmjb"),
	mEmail.Gmail("mo7trade2@gmail.com", "mhaqiyalgaiyhoto"),
	mEmail.Gmail("meichangliang@gmail.com", "pwlooxzamplnwwgf"),
	mEmail.QQ("meichangliang@qq.com", "fxdxnbyronppbfha"),
	mEmail.QQ("mo7trade@qq.com", "aluanmhgxubnbigf"),
}

// ======构建邮件的封装===========
type EmailOpt struct {
	From     string
	To       []string
	Subject  string
	TmplName string
	SendData any
}

func BuildEmail(opt EmailOpt) *mEmail.EmailInfo {
	if len(opt.From) < 1 {
		opt.From = "Message.net"
	}
	if len(opt.To) < 1 {
		opt.To = []string{"trade@mo7.cc"}
	}
	if len(opt.Subject) < 1 {
		opt.Subject = "Default Subject"
	}

	TemplateStr := tmpl.SysEmail

	switch opt.TmplName {
	case "SysEmail":
		TemplateStr = tmpl.SysEmail
	case "CodeEmail":
		TemplateStr = tmpl.CodeEmail
	case "RegisterEmail":
		TemplateStr = tmpl.RegisterEmail
	}

	EmailServe := EmailAccountList[0]

	mJson.Println(EmailServe)

	emailObj := mEmail.New(mEmail.Opt{
		Account:     EmailServe.Account,
		Password:    EmailServe.Password,
		Port:        EmailServe.Port,
		Host:        EmailServe.Host,
		To:          opt.To,
		From:        opt.From,
		Subject:     opt.Subject,
		TemplateStr: TemplateStr,
		SendData:    opt.SendData,
	})
	return emailObj
}
