package await

import (
	"Message.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/gofiber/fiber/v2"
)

/*
/api/await
*/

func Router(router fiber.Router) {
	r := router.Group("/await", MiddleWare)

	r.Post("/xxx", xxx)
}

func xxx(c *fiber.Ctx) error {
	json := mFiber.Parser(c)

	return c.JSON(result.Succeed.WithData(json))
}
