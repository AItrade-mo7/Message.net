package api

import (
	"Message.net/server/global/config"
	"Message.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func GetSysEmailList(c *fiber.Ctx) error {
	EmailList := []string{}

	for _, val := range config.MyEmailList {
		EmailList = append(EmailList, val.Account)
	}

	return c.JSON(result.Succeed.WithData(EmailList))
}
