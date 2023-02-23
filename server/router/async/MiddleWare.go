package async

import (
	"Message.net/server/global/config"
	"github.com/gofiber/fiber/v2"
)

// 异步任务
func MiddleWare(c *fiber.Ctx) error {
	c.Set("Data-Path", config.SysName+"/api/async")

	return c.Next()
}
