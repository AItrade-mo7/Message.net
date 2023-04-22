package global

import (
	"strings"

	"Message.net/server/global/config"
	"Message.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type EmailCountType struct {
	Hour   int
	Hour24 int
}

var EmailCount map[string]EmailCountType

func StoreSendEmail(storeOpt dbType.MessageEmail) {
	db, err := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Message",
	}).Connect()
	if err != nil {
		LogErr("disposeTask.StoreSendEmail", err)
		return
	}
	defer Run.Println("global.StoreSendEmail 关闭数据库", storeOpt.EmailID)
	defer db.Close()
	db.Collection("Email")

	db.Table.InsertOne(db.Ctx, storeOpt)
}

func SendEmail(emailOpt mEmail.Opt) error {
	err := mEmail.New(emailOpt).Send()
	storeJsonByte := mJson.ToJson(emailOpt)
	var StoreOpt dbType.MessageEmail
	jsoniter.Unmarshal(storeJsonByte, &StoreOpt)
	StoreOpt.SendResult = mStr.ToStr(err)
	StoreOpt.EmailID = mEncrypt.GetUUID()
	StoreOpt.CreateTime = mTime.GetUnixInt64()
	StoreOpt.CreateTimeStr = mTime.UnixFormat(StoreOpt.CreateTime)
	Run.Println("邮件已发送", StoreOpt.EmailID, StoreOpt.Subject, StoreOpt.To, err)
	StoreSendEmail(StoreOpt) // 存储发送记录
	return err
}

// ======== 账号池子 ==============

func GetEmailServe() (resData mEmail.ServeType) {
	Len := len(config.MyEmailList)
	index := mCount.GetRound(0, int64(Len-1))

	resData = config.MyEmailList[index]
	HourCount := EmailCount[resData.Account]

	isQQ := strings.Contains(resData.Account, "qq.com")
	isGmail := strings.Contains(resData.Account, "gmail.com")

	if isQQ { // QQ 邮箱 发件次数 和 频率 超过
		if HourCount.Hour > 20 && HourCount.Hour24 > 100 {
			return GetEmailServe()
		}
	}

	if isGmail { // Gmail 邮箱 发件次数 和 频率 超过
		if HourCount.Hour > 20 && HourCount.Hour24 > 100 {
			return GetEmailServe()
		}
	}

	return
}
