package public

import (
	"Message.net/server/router/api"
	"github.com/gofiber/fiber/v2"
)

/*
/api/public
*/

func Router(router fiber.Router) {


	r := router.Group("/public", MiddleWare)

	r.Post("/InsertTaskQueue", api.InsertTaskQueue)
}
