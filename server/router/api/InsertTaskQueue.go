package api

import (
	"fmt"

	"Message.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

func InsertTaskQueue(c *fiber.Ctx) error {
	var json map[string]any
	mFiber.Parser(c, &json)

	fmt.Println(json)

	return c.JSON(result.Succeed.WithData("Succeed"))
}
