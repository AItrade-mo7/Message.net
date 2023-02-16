package async

import (
	"github.com/gofiber/fiber/v2"
)

func MiddleWare(c *fiber.Ctx) error {
	c.Set("Data-Path", "Message.net/api/async")

	return c.Next()
}
