package disposeTask

import "fmt"

func SendSysEmail() {
	fmt.Println("发送系统邮件")
}

func SendCodeEmail() {
	fmt.Println("发送验证码")
}

func SendRegisterEmail() {
	fmt.Println("发送注册通知")
}

// 发送邮件的
// 用 gmail 发送
func FromGmail() {
	/*
		meichangliang@gmail.com
		nmqlusfgaeyexxok
	*/
}

// 用 企业微信 发送
func FromWorkWeiXin() {
	/*
		trade@mo7.cc
		DXir4WLb2aGaknLZ
	*/
}

// 用 qq 发送
func FromQQ() {
	/*
		670188307@qq.com
		meichangliang@qq.com
		fxdxnbyronppbfha

	*/
}
