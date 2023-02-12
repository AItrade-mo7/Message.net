package global

import (
	"fmt"

	"Message.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

func StoreSendEmail(emailOpt dbType.MessageEmail) {
	fmt.Println("数据库存储")

	mJson.Println(emailOpt)
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
