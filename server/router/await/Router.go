package await

import (
	"github.com/gofiber/fiber/v2"
)

/*
/api/await
*/

func Router(router fiber.Router) {
	r := router.Group("/await", MiddleWare)

	r.Post("/xxx", xxx)
}
