package disposeTask

import (
	"fmt"
	"strings"

	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTask"
	jsoniter "github.com/json-iterator/go"
)

var EmailAccountList = []mEmail.ServeType{
	mEmail.Gmail("meichangliang@gmail.com", "nmqlusfgaeyexxok"),
	mEmail.Gmail("mo7trade1@gmail.com", "asdasd55555"),
	mEmail.Gmail("mo7trade2@gmail.com", "asdasd55555"),
	mEmail.QQ("meichangliang@qq.com", "fxdxnbyronppbfha"),
	mEmail.WorkWeiXin("trade@mo7.cc", "DXir4WLb2aGaknLZ"),
}

// 发送邮件的函数
type EmailOpt struct {
	From     string
	To       []string
	Subject  string
	Template string
	SendData any
}

var EmailTmplList = []string{
	"SysEmail",
	"CodeEmail",
	"RegisterEmail",
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

	sepChart := "&"
	TmplAll := strings.Join(EmailTmplList, sepChart)
	TmplAll = mStr.Join(TmplAll, sepChart)
	TmplNow := mStr.Join(opt.Template, sepChart)
	findTmpl := strings.Contains(TmplAll, TmplNow)

	if !findTmpl || TmplNow == sepChart {
		opt.Template = EmailTmplList[0]
	}

	mJson.Println(opt)

	emailObj := mEmail.New(mEmail.Opt{
		To:          opt.To,
		From:        opt.From,
		Subject:     opt.Subject,
		TemplateStr: opt.Template,
		SendData:    opt.SendData,
	})
	return emailObj
}

// 发送完邮件之后要把发送记录存储到数据库中去
func SendSysEmail(opt any) {
	jsonByte := mJson.ToJson(opt)
	var info mTask.SendEmail
	jsoniter.Unmarshal(jsonByte, &info)

	emailObj := BuildEmail(EmailOpt{
		From:     info.From,
		To:       info.To,
		Subject:  info.Subject,
		Template: info.TmplName,
		SendData: info.SendData,
	})

	err := emailObj.Send()
	fmt.Println("邮件发送完成", err)
}
