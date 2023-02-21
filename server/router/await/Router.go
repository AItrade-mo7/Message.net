package await

import (
	"Message.net/server/router/api"
	"github.com/gofiber/fiber/v2"
)

/*
/api/await
*/

func Router(router fiber.Router) {
	r := router.Group("/await", MiddleWare)

	r.Post("/SendEmailCode", api.SendEmailCode)
}
