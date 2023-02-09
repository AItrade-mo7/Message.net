package global

import (
	"log"

	"Message.net/server/global/config"
	"github.com/EasyGolang/goTools/mLog"
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
}
