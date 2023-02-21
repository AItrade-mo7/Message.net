package config

import "github.com/EasyGolang/goTools/mEmail"

var MyEmailList = []mEmail.ServeType{
	// gmail
	mEmail.Gmail("mo7trade1@gmail.com", "bhmfbovjxnkmcmjb"),
	mEmail.Gmail("mo7trade2@gmail.com", "mhaqiyalgaiyhoto"),
	mEmail.Gmail("mo7trade3@gmail.com", "oinleqxgslppsfbx"),
	mEmail.Gmail("meichangliang@gmail.com", "pwlooxzamplnwwgf"),
	// qq
	mEmail.QQ("mo7trade@qq.com", "aluanmhgxubnbigf"),
	mEmail.QQ("meichangliang@qq.com", "fxdxnbyronppbfha"),
}

var SysEmail = "meichangliang@outlook.com"
