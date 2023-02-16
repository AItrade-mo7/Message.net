package disposeTask

import (
	"Message.net/server/global"
	"Message.net/server/global/config"
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
	if err != nil {
		//
		if info.TmplName == "CodeEmail" {
			go EmailAction(info)
		}
	}

	return err
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
		opt.From = config.SysName
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

	EmailServe := global.GetEmailServe()

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

// 邮件任务的后续处理
func EmailAction(info mTask.SendEmail) {
	// 判断是否为 验证码
	if info.TmplName == "CodeEmail" {
		jsonByte := mJson.ToJson(info.SendData)
		var SendData mTask.CodeEmailParam
		jsoniter.Unmarshal(jsonByte, &SendData)

		if len(SendData.VerifyCode) > 0 && len(info.To) > 0 {
			UpdateEmailCode(info.To, SendData.VerifyCode)
		} else {
			global.LogErr("disposeTask.EmailAction 空验证码", mJson.Format(info))
		}
	}
}
