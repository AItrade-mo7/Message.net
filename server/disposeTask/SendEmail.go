package disposeTask

import (
	"Message.net/server/global"
	"Message.net/server/tmpl"
	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTask"
	jsoniter "github.com/json-iterator/go"
)

func SendSysEmail(opt any) error {
	jsonByte := mJson.ToJson(opt)
	var info mTask.SendEmail
	jsoniter.Unmarshal(jsonByte, &info)

	emailOpt := GetEmailOpt(EmailOpt{
		From:     info.From,
		To:       info.To,
		Subject:  info.Subject,
		TmplName: info.TmplName,
		SendData: info.SendData,
	})

	err := global.SendEmail(emailOpt) // 发送并存储记录
	return err
}

// ======== 账号池子 ==============
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

func GetEmailOpt(opt EmailOpt) mEmail.Opt {
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

	emailOpt := mEmail.Opt{
		Account:     EmailServe.Account,
		Password:    EmailServe.Password,
		Port:        EmailServe.Port,
		Host:        EmailServe.Host,
		To:          opt.To,
		From:        opt.From,
		Subject:     opt.Subject,
		TemplateStr: TemplateStr,
		SendData:    opt.SendData,
	}
	return emailOpt
}
