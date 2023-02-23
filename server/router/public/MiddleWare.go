package public

import (
	"Message.net/server/global/config"
	"github.com/gofiber/fiber/v2"
)

func MiddleWare(c *fiber.Ctx) error {
	c.Set("Data-Path", config.SysName+"/api/await")

	return c.Next()
}
