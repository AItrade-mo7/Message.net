package ready

import (
	"Message.net/server/global/config"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
)

func StartEmail() {
	TaskContent := mJson.StructToMap(mTask.SysEmail{
		From:    config.SysName,
		To:      []string{config.SysEmail},
		Subject: "服务启动",
		SendData: mTask.SysEmailParam{
			Title:        "Message.net 服务启动",
			Message:      "服务启动",
			Content:      mJson.Format(config.AppInfo),
			SysTime:      mTime.UnixFormat(mTime.GetUnix()),
			Source:       config.SysName,
			SecurityCode: "trade.mo7.cc",
		},
	})

	now := mTime.GetTime()
	NewTask := mTask.TaskType{
		TaskID:        mEncrypt.GetUUID(),
		TaskType:      "SysEmail",
		Content:       TaskContent,
		Source:        config.SysName,
		Description:   config.SysName + "程序启动",
		CreateTime:    now.TimeUnix,
		CreateTimeStr: now.TimeStr,
	}
	FilePath := mStr.Join(
		config.Dir.TaskQueue,
		"/",
		NewTask.TaskID+".json",
	)
	mFile.Write(FilePath, mJson.ToStr(NewTask))
}
