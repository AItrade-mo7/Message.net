package disposeTask

import (
	"Message.net/server/global"
	"Message.net/server/tmpl"
	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTask"
)

// ======= 发送系统邮件 =======
func SendSysEmail(TaskCont mTask.SysEmail) error {
	EmailServe := global.GetEmailServe()
	emailOpt := mEmail.Opt{
		Account:     EmailServe.Account,
		Password:    EmailServe.Password,
		Port:        EmailServe.Port,
		Host:        EmailServe.Host,
		To:          TaskCont.To,
		From:        TaskCont.From,
		Subject:     TaskCont.Subject,
		TemplateStr: tmpl.SysEmail, //  采用系统模板
		SendData:    TaskCont.SendData,
	}
	// 发送并存储记录
	err := global.SendEmail(emailOpt)
	return err
}

// ======= 发送 验证码 邮件 =======
func SendCodeEmail(TaskCont mTask.CodeEmail) error {
	EmailServe := global.GetEmailServe()
	emailOpt := mEmail.Opt{
		Account:     EmailServe.Account,
		Password:    EmailServe.Password,
		Port:        EmailServe.Port,
		Host:        EmailServe.Host,
		To:          TaskCont.To,
		From:        TaskCont.From,
		Subject:     TaskCont.Subject,
		TemplateStr: tmpl.CodeEmail, //  采用验证码模板
		SendData:    TaskCont.SendData,
	}
	// 发送并存储记录
	err := global.SendEmail(emailOpt)
	// err 为 nil 的时候
	if err == nil {
		if len(TaskCont.SendData.VerifyCode) > 0 && len(TaskCont.To) > 0 {
			UpdateEmailCode(TaskCont)
		} else {
			global.LogErr("disposeTask.EmailAction 空验证码", mJson.Format(TaskCont))
		}
	}

	return err
}

// ======= 发送注册成功邮件 =======
func SendRegisterSucceedEmail(TaskCont mTask.RegisterSucceedEmail) error {
	EmailServe := global.GetEmailServe()
	emailOpt := mEmail.Opt{
		Account:     EmailServe.Account,
		Password:    EmailServe.Password,
		Port:        EmailServe.Port,
		Host:        EmailServe.Host,
		To:          TaskCont.To,
		From:        TaskCont.From,
		Subject:     TaskCont.Subject,
		TemplateStr: tmpl.RegisterSucceedEmail, //  采用系统模板
		SendData:    TaskCont.SendData,
	}
	// 发送并存储记录
	err := global.SendEmail(emailOpt)
	return err
}
