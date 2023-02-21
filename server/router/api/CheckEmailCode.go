package api

import (
	"fmt"

	"Message.net/server/global/config"
	"Message.net/server/global/dbType"
	"Message.net/server/router/result"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CheckEmailCodeParam struct {
	Email string
	Code  string // 加密的
}

func CheckEmailCode(c *fiber.Ctx) error {
	var json CheckEmailCodeParam
	mFiber.Parser(c, &json)

	isEmail := mVerify.IsEmail(json.Email)
	if !isEmail {
		emailErr := fmt.Errorf("json.Email 格式不正确 %+v", json.Email)
		return c.JSON(result.ErrEmailCode.WithMsg(emailErr))
	}

	if len(json.Code) < 1 {
		emailErr := fmt.Errorf("json.Code 不能为空")
		return c.JSON(result.ErrEmailCode.WithMsg(emailErr))
	}

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Message",
	}).Connect().Collection("VerifyCode")
	defer db.Close()

	// 查找参数设置
	FK := bson.D{{
		Key:   "Email",
		Value: json.Email,
	}}
	findOptions := options.FindOne()
	// 查找参数设置

	var dbData dbType.EmailCodeTable
	db.Table.FindOne(db.Ctx, FK, findOptions).Decode(&dbData)

	// 校验时间
	sendTime := mStr.ToStr(dbData.SendTime)
	nowTime := mTime.GetUnix()
	subStr := mCount.Sub(nowTime, sendTime)

	if mCount.Le(subStr, mCount.Mul(mTime.UnixTime.Minute, "12")) > 0 {
		err := fmt.Errorf("验证码已过期")
		return c.JSON(result.ErrEmailCode.WithMsg(err))
	}

	DBCode := mEncrypt.MD5(dbData.Code)
	if DBCode != json.Code {
		err := fmt.Errorf("验证码不正确")
		return c.JSON(result.ErrEmailCode.WithMsg(err))
	}

	return c.JSON(result.Succeed.WithMsg("Succeed"))
}
