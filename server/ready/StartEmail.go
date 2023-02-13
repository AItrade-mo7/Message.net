package ready

import (
	"Message.net/server/global"
	"Message.net/server/global/config"
	"Message.net/server/tmpl"
	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
)

func StartEmail() {
	EmailServe := mEmail.Gmail("meichangliang@gmail.com", "pwlooxzamplnwwgf")
	emailOpt := mEmail.Opt{
		Account:  EmailServe.Account,
		Password: EmailServe.Password,
		Port:     EmailServe.Port,
		Host:     EmailServe.Host,
		To: []string{
			"trade@mo7.cc",
		},
		From:        "Message.net",
		Subject:     "服务启动",
		TemplateStr: tmpl.SysEmail,
		SendData: mTask.SysEmailParam{
			Title:        "Message.net 服务启动",
			Message:      "服务启动",
			Content:      mJson.Format(config.AppInfo),
			SysTime:      mTime.UnixFormat(mTime.GetUnix()),
			Source:       "Message.net",
			SecurityCode: "trade.mo7.cc",
		},
	}
	global.SendEmail(emailOpt)
}