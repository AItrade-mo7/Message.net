package api

import (
	"fmt"

	"Message.net/server/global/config"
	"Message.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTask"
	"github.com/gofiber/fiber/v2"
)

func InsertTaskQueue(c *fiber.Ctx) error {
	var json mTask.TaskType
	mFiber.Parser(c, &json)

	nowTaskType, err := mTask.CheckTask(json)
	if err != nil {
		return c.JSON(result.Fail.WithMsg(err))
	}
	// 把任务写到 目录当中
	FilePath := mStr.Join(
		config.Dir.TaskQueue,
		"/",
		nowTaskType.TaskID+".json",
	)

	fmt.Println(FilePath)

	mFile.Write(FilePath, mJson.ToStr(json))

	return c.JSON(result.Succeed.WithData("Succeed"))
}
