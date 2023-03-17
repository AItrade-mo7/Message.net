package disposeTask

import (
	"os"
	"path/filepath"

	"Message.net/server/global"
	"Message.net/server/global/config"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

func Treatment() {
	global.Run.Println("disposeTask.Treatment 执行一次目录遍历")
	// 在这里连接数据库
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Message",
	}).Connect().Collection("Task")
	defer global.Run.Println("disposeTask.StoreTask 关闭数据库")
	defer db.Close()

	fsList, err := os.ReadDir(config.Dir.TaskQueue)
	if err != nil {
		// 错误处理
		global.LogErr("disposeTask.Treatment", err)
	}
	for _, file := range fsList {
		if file.IsDir() {
			continue
		}
		Path := filepath.Join(config.Dir.TaskQueue, file.Name())
		isPath := mPath.Exists(Path)
		if isPath {
			ReadTask(Path, db)
		}
	}
}

func ReadTask(path string, db *mMongo.DB) {
	json, err := os.ReadFile(path)
	if err != nil {
		global.LogErr("disposeTask.ReadTask-1", err)
	}

	var Task mTask.TaskType
	jsoniter.Unmarshal(json, &Task)

	// 任务分配执行区

	switch Task.TaskType {
	case "SysEmail":
		var SysObj mTask.SysEmail
		jsoniter.Unmarshal(mJson.ToJson(Task.Content), &SysObj)
		err = SendSysEmail(SysObj)
	case "CodeEmail":
		var CodeObj mTask.CodeEmail
		jsoniter.Unmarshal(mJson.ToJson(Task.Content), &CodeObj)
		err = SendCodeEmail(CodeObj)
	case "RegisterSucceedEmail":
		var RegisterSucceedObj mTask.RegisterSucceedEmail
		jsoniter.Unmarshal(mJson.ToJson(Task.Content), &RegisterSucceedObj)
		err = SendRegisterSucceedEmail(RegisterSucceedObj)
	}

	// 任务存储
	Task.EndTime = mTime.GetUnixInt64()
	Task.EndTimeStr = mTime.UnixFormat(Task.EndTime)
	Task.Result = mStr.ToStr(err)

	db.Table.InsertOne(db.Ctx, Task)

	global.Run.Println("===== 一条 disposeTask.ReadTask 任务执行结束 ======", Task.TaskID)
	if err != nil {
		global.LogErr("disposeTask.ReadTask-2 任务执行失败!", Task.TaskID)
		return
	} else {
		err = os.Remove(path)
		if err != nil {
			global.LogErr("disposeTask.ReadTask-3 任务删除失败!", err)
		}
	}
}
