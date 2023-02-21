package testCase

import (
	"fmt"

	"Message.net/server/global"
	"Message.net/server/global/config"
	"Message.net/server/tmpl"
	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
)

// ======== 测试Demo ==========
func SendEmail() {
	for _, val := range config.MyEmailList {
		EmailServe := val
		emailOpt := mEmail.Opt{
			Account:     EmailServe.Account,
			Password:    EmailServe.Password,
			Port:        EmailServe.Port,
			Host:        EmailServe.Host,
			To:          []string{config.SysEmail},
			From:        config.SysName,
			Subject:     "测试邮件发送",
			TemplateStr: tmpl.SysEmail,
			SendData: mTask.SysEmailParam{
				Title:        "测试邮件发送",
				Message:      "当前邮箱账号为: " + EmailServe.Account,
				Content:      mJson.Format(EmailServe),
				SysTime:      mTime.UnixFormat(mTime.GetUnix()),
				Source:       config.SysName,
				SecurityCode: "trade.mo7.cc",
			},
		}
		err := global.SendEmail(emailOpt)
		fmt.Println(EmailServe.Account, err)
	}
}
