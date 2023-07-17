package main

import (
	_ "embed"

	"Message.net/server/global"
	"Message.net/server/global/config"
	"Message.net/server/ready"
	"Message.net/server/testCase"
	jsoniter "github.com/json-iterator/go"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)

	// 初始化系统参数
	global.Start()

	// // 主要业务逻辑准备
	ready.Start()

	// // 启动 http 监听服务
	// router.Start()

	testCase.SendEmail()
}
