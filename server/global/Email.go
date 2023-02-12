package global

import (
	"Message.net/server/global/config"
	"Message.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

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

	Run.Println("邮件已发送", StoreOpt.EmailID, StoreOpt.Subject, StoreOpt.To, err)

	StoreSendEmail(StoreOpt) // 存储发送记录

	return err
}