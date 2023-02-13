package disposeTask

import (
	"Message.net/server/global"
	"Message.net/server/tmpl"
	"github.com/EasyGolang/goTools/mCount"
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
	mEmail.QQ("mo7trade@qq.com", "aluanmhgxubnbigf"),
	mEmail.QQ("meichangliang@qq.com", "fxdxnbyronppbfha"),
	mEmail.Gmail("meichangliang@gmail.com", "pwlooxzamplnwwgf"),
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

	EmailServe := GetEmailServe()

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

func GetEmailServe() (resData mEmail.ServeType) {
	Len := len(EmailAccountList)
	index := mCount.GetRound(0, int64(Len-1))

	resData = EmailAccountList[index]
	HourCount := global.UseEmailCountHour[resData.Account]
	Hour24Count := global.UseEmailCount24Hour[resData.Account]

	// 小时内 20 封 ; 24 小时内 100 封
	if HourCount < 20 && Hour24Count < 100 {
		return
	}
	return GetEmailServe()
}
