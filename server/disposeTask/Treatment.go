package disposeTask

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"Message.net/server/global/config"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mTask"
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
		SendSysEmail(Task.Content)
	}

	// 任务处理完要把任务存储到数据库当中去
}
