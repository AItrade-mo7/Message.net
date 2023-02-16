package await

import (
	"Message.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

func xxx(c *fiber.Ctx) error {
	json := mFiber.Parser(c)

	return c.JSON(result.Succeed.WithData(json))
}
