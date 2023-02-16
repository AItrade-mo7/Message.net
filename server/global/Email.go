package global

import (
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
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Message",
	}).Connect().Collection("Email")
	defer Run.Println("global.StoreSendEmail 关闭数据库", storeOpt.EmailID)
	defer db.Close()

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
var EmailAccountList = []mEmail.ServeType{
	mEmail.Gmail("mo7trade1@gmail.com", "bhmfbovjxnkmcmjb"),
	mEmail.Gmail("mo7trade2@gmail.com", "mhaqiyalgaiyhoto"),
	mEmail.QQ("mo7trade@qq.com", "aluanmhgxubnbigf"),
	mEmail.QQ("meichangliang@qq.com", "fxdxnbyronppbfha"),
	mEmail.Gmail("meichangliang@gmail.com", "pwlooxzamplnwwgf"),
}

func GetEmailServe() (resData mEmail.ServeType) {
	Len := len(EmailAccountList)
	index := mCount.GetRound(0, int64(Len-1))

	resData = EmailAccountList[index]
	HourCount := EmailCount[resData.Account]

	// 小时内 20 封 ; 24 小时内 100 封
	if HourCount.Hour < 20 && HourCount.Hour24 < 100 {
		return
	}
	return GetEmailServe()
}
