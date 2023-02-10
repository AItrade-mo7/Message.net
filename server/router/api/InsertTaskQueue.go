package api

import (
	"fmt"

	"Message.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

type TaskQueue struct {
	TaskID   string         `bson:"TaskID"`   // 任务ID
	TaskType string         `bson:"TaskType"` // 任务类型  SendEmail, SendOrder 等
	Content  map[string]any `bson:"Content"`  // 任务内容 用不同的模板去解析
}

func InsertTaskQueue(c *fiber.Ctx) error {
	var json map[string]any
	mFiber.Parser(c, &json)

	fmt.Println(json)

	return c.JSON(result.Succeed.WithData("Succeed"))
}
