package async

import (
	"github.com/gofiber/fiber/v2"
)

/*
/api/async
*/

func Router(router fiber.Router) {
	r := router.Group("/async", MiddleWare)

	r.Post("/InsertTaskQueue", InsertTaskQueue)
}
