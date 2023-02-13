package global

import (
	"fmt"
	"log"

	"Message.net/server/global/config"
	"Message.net/server/tmpl"
	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mLog"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
)

var (
	Log *log.Logger // 系统日志& 重大错误或者事件
	Run *log.Logger //  运行日志
)

func LogInit() {
	// 创建一个log
	Log = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Sys",
	})

	Run = mLog.NewLog(mLog.NewLogParam{
		Path: config.Dir.Log,
		Name: "Run",
	})

	// 设定清除 log
	mLog.Clear(mLog.ClearParam{
		Path:      config.Dir.Log,
		ClearTime: mTime.UnixTimeInt64.Day * 10,
	})

	config.LogErr = LogErr
	config.Log = Log
}

func LogErr(sum ...any) {
	str := fmt.Sprintf("系统错误: %+v", sum)
	Log.Println(str)

	// 系统的重大错误，必须要发送错误邮件 默认使用企业微信发送
	EmailServe := mEmail.Gmail("meichangliang@gmail.com", "pwlooxzamplnwwgf")
	message := ""
	if len(sum) > 0 {
		message = mStr.ToStr(sum[0])
	}
	content := mJson.Format(sum)
	emailOpt := mEmail.Opt{
		Account:     EmailServe.Account,
		Password:    EmailServe.Password,
		Port:        EmailServe.Port,
		Host:        EmailServe.Host,
		To:          []string{"trade@mo7.cc"},
		From:        "Message.net",
		Subject:     "系统错误",
		TemplateStr: tmpl.SysEmail,
		SendData: mTask.SysEmailParam{
			Title:        "Message.net 系统出错",
			Message:      message,
			Content:      content,
			SysTime:      mTime.UnixFormat(mTime.GetUnix()),
			Source:       "Message.net",
			SecurityCode: "trade.mo7.cc",
		},
	}

	err := SendEmail(emailOpt)
	Log.Println("错误邮件已发", err)
}
