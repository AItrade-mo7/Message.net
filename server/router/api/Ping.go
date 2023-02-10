package api

import (
	"Message.net/server/global/config"
	"Message.net/server/router/result"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type AppInfoType struct {
	Name    string `bson:"name"`
	Version string `bson:"version"`
}

func Ping(c *fiber.Ctx) error {
	json := mFiber.Parser(c)

	var ApiInfo AppInfoType
	jsoniter.Unmarshal(mJson.ToJson(config.AppInfo), &ApiInfo)

	ReturnData := make(map[string]any)
	ReturnData["ResParam"] = json
	ReturnData["Method"] = c.Method()
	ReturnData["ApiInfo"] = ApiInfo

	ReturnData["UserAgent"] = c.Get("User-Agent")
	ReturnData["Path"] = c.OriginalURL()

	ips := c.IPs()
	if len(ips) > 0 {
		ReturnData["IP"] = ips[0]
	}

	return c.JSON(result.Succeed.WithData(ReturnData))
}
