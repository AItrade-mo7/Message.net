package config

import "github.com/EasyGolang/goTools/mEmail"

var MyEmailList = []mEmail.ServeType{
	// gmail
	mEmail.Gmail("1234@gmail.com", "1234"),
}

var SysEmail = "meichangliang@outlook.com"
