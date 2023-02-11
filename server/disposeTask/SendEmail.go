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

		Account:  "meichangliang@gmail.com",
		Password: "nmqlusfgaeyexxok",
		Port:     "587",
		Host:     "smtp.gmail.com",

	*/
}

// 用 企业微信 发送
func FromWorkWeiXin() {
	/*
		Account:  "trade@mo7.cc",
		Password: "DXir4WLb2aGaknLZ",
		Port:     "587",
		Host:     "smtp.exmail.qq.com",
	*/
}

// 用 qq 发送
func FromQQ() {
	/*
		Account: "meichangliang@qq.com",
		// Account:  "670188307@qq.com",
		Password: "fxdxnbyronppbfha",
		Port:     "587",
		Host:     "smtp.qq.com",
	*/
}
