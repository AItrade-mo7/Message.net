package middle

import (
	"errors"
	"strings"

	"Message.net/server/global/config"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
)

func EncryptAuth(c *fiber.Ctx) error {
	EncStr := c.Get("Auth-Encrypt")
	if len(EncStr) < 20 {
		return errors.New("需要授权码")
	}

	shaStr := config.Encrypt(c.Path() + c.Get("User-Agent") + mStr.ToStr(c.Body()))
	isFind := strings.Contains(shaStr, EncStr)

	if !isFind {
		return errors.New("授权验证错误")
	}

	return nil
}
