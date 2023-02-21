package api

import (
	"fmt"

	"Message.net/server/global"
	"Message.net/server/global/config"
	"Message.net/server/global/dbType"
	"Message.net/server/router/result"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFiber"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/EasyGolang/goTools/mVerify"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SendEmailParam struct {
	Email  string
	Action string
}

func SendEmailCode(c *fiber.Ctx) error {
	var json SendEmailParam
	mFiber.Parser(c, &json)

	isEmail := mVerify.IsEmail(json.Email)
	if !isEmail {
		emailErr := fmt.Errorf("json.Email 格式不正确 %+v", json.Email)
		return c.JSON(result.ErrEmail.WithMsg(emailErr))
	}

	if len(json.Action) < 1 {
		emailErr := fmt.Errorf("json.Action 不能为空")
		return c.JSON(result.ErrEmail.WithMsg(emailErr))
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

	sendTime := mStr.ToStr(dbData.SendTime) // 发送验证码的时间
	nowTime := mTime.GetUnix()              // 当前时间
	subStr := mCount.Sub(nowTime, sendTime) // 两者的时间差

	// 逻辑，如果没有发送则 sendTime 为 0 ， 则时间差无限大，如果发送了
	if mCount.Le(subStr, mCount.Mul(mTime.UnixTime.Minute, "10")) < 0 {
		db.Close()
		err := fmt.Errorf("刚刚的验证码还能再用哦~")
		return c.JSON(result.Fail.WithMsg(err))
	}

	// 创建异步的验证码发送任务
	TaskContent := mJson.StructToMap(mTask.CodeEmail{
		From:    "AItrade",
		To:      json.Email,
		Subject: "请查收您的验证码",
		SendData: mTask.CodeEmailParam{
			VerifyCode:   mVerify.NewCode(),
			Action:       json.Action,
			SysTime:      mTime.GetTime().TimeStr,
			Source:       config.SysName,
			SecurityCode: "trade.mo7.cc",
		},
	})

	now := mTime.GetTime()
	NewTask := mTask.TaskType{
		TaskID:        mEncrypt.GetUUID(),
		TaskType:      "CodeEmail",
		Content:       TaskContent,
		Source:        config.SysName,
		Description:   "验证码邮件", // 任务描述
		CreateTime:    now.TimeUnix,
		CreateTimeStr: now.TimeStr,
	}

	FilePath := mStr.Join(
		config.Dir.TaskQueue,
		"/",
		NewTask.TaskID+".json",
	)
	mFile.Write(FilePath, mJson.ToStr(json))
	global.TaskChan <- NewTask.TaskID

	return c.JSON(result.Succeed.WithMsg("验证码已发送"))
}
