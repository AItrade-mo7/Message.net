package disposeTask

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"Message.net/server/global"
	"Message.net/server/global/config"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

func Treatment() {
	fmt.Println("开始遍历任务目录")
	fsList, err := os.ReadDir(config.Dir.TaskQueue)
	if err != nil {
		// 错误处理
		fmt.Println("ready.Treatment", err)
	}
	for _, file := range fsList {
		if file.IsDir() {
			continue
		}
		Path := filepath.Join(config.Dir.TaskQueue, file.Name())
		isPath := mPath.Exists(Path)
		if isPath {
			ReadTask(Path)
		}
	}
}

func ReadTask(path string) {
	json, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var Task mTask.TaskType
	jsoniter.Unmarshal(json, &Task)

	switch Task.TaskType {
	case "SendEmail":
		err = SendSysEmail(Task.Content)
	}

	Task.EndTime = mTime.GetUnixInt64()
	Task.EndTimeStr = mTime.UnixFormat(Task.EndTime)
	Task.Result = mStr.ToStr(err)
	StoreTask(Task)
}

func StoreTask(task mTask.TaskType) {
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Message",
	}).Connect().Collection("Task")
	defer global.Run.Println("disposeTask.StoreTask 关闭数据库", task.TaskID)
	defer db.Close()
	db.Table.InsertOne(db.Ctx, task)
}
