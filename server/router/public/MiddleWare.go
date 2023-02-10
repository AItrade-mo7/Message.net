package public

import (
	"github.com/gofiber/fiber/v2"
)

func MiddleWare(c *fiber.Ctx) error {
	c.Set("Data-Path", "Message.net/api/public")

	return c.Next()
}
