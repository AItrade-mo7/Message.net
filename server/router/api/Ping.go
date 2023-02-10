package api

import (
	"Message.net/server/global/config"
	"Message.net/server/router/result"
	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type AppInfoType struct {
	Name    string `bson:"name"`
	Version string `bson:"version"`
}

func Ping(c *fiber.Ctx) error {
	json := mFiber.Parser(c)

	// 在这里请求数据
	ClientFileReqData, _ := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: "http://trade.mo7.cc",
		Path:   "/package.json?tmp=" + mTime.GetUnix(),
	}).Get()

	var ClientInfo AppInfoType
	jsoniter.Unmarshal(ClientFileReqData, &ClientInfo)

	var ApiInfo AppInfoType
	jsoniter.Unmarshal(mJson.ToJson(config.AppInfo), &ApiInfo)

	ReturnData := make(map[string]any)
	ReturnData["ResParam"] = json
	ReturnData["Method"] = c.Method()
	ReturnData["ApiInfo"] = ApiInfo
	ReturnData["ClientInfo"] = ClientInfo

	ReturnData["UserAgent"] = c.Get("User-Agent")
	ReturnData["Path"] = c.OriginalURL()

	ips := c.IPs()
	if len(ips) > 0 {
		ReturnData["IP"] = ips[0]
	}

	return c.JSON(result.Succeed.WithData(ReturnData))
}
