package config

import (
	"os"

	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
)

var Dir struct {
	Home      string // Home 根目录
	App       string // APP 根目录
	Log       string // 日志文件目录
	JsonData  string // json 数据存放目录
	TaskQueue string // 任务队列
}

var File struct {
	SysEnv      string // ~/sys_env.yaml
	LocalSysEnv string // ./sys_env.yaml
}

func DirInit() {
	Dir.Home = mPath.HomePath()

	Dir.App, _ = os.Getwd()

	Dir.Log = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"logs",
	)

	Dir.JsonData = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"jsonData",
	)
	Dir.TaskQueue = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"TaskQueue",
	)

	File.SysEnv = mStr.Join(
		Dir.Home,
		mStr.ToStr(os.PathSeparator),
		"sys_env.yaml",
	)
	File.LocalSysEnv = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"sys_env.yaml",
	)

	// 检测 JsonData 目录
	isJsonDataPath := mPath.Exists(Dir.JsonData)
	if !isJsonDataPath {
		// 不存在则创建 logs 目录
		os.MkdirAll(Dir.JsonData, 0o777)
	}

	// 检测 logs 目录
	isLogPath := mPath.Exists(Dir.Log)
	if !isLogPath {
		// 不存在则创建 logs 目录
		os.MkdirAll(Dir.Log, 0o777)
	}

	// 检测 TaskQueue 目录
	isTaskQueuePath := mPath.Exists(Dir.TaskQueue)
	if !isTaskQueuePath {
		// 不存在则创建 logs 目录
		os.MkdirAll(Dir.TaskQueue, 0o777)
	}
}
