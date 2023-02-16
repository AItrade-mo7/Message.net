package disposeTask

import (
	"Message.net/server/global"
	"Message.net/server/global/config"
	"Message.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStruct"
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateEmailCode(info mTask.SendEmail) {
	jsonByte := mJson.ToJson(info.SendData)
	var SendData mTask.CodeEmailParam
	jsoniter.Unmarshal(jsonByte, &SendData)
	mJson.Println(info)

	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Message",
	}).Connect().Collection("VerifyCode")
	defer global.Run.Println("disposeTask.UpdateEmailCode 关闭数据库")
	defer db.Close()

	FK := bson.D{}
	UK := bson.D{}

	nowTime := mTime.GetUnixInt64()

	for _, val := range info.To {
		FK = append(FK, bson.E{
			Key:   "Email",
			Value: val,
		})
		EmailCode := dbType.EmailCodeTable{
			Email:    val,
			Code:     SendData.VerifyCode,
			SendTime: nowTime,
		}
		mStruct.Traverse(EmailCode, func(key string, val any) {
			UK = append(UK, bson.E{
				Key: "$set",
				Value: bson.D{
					{
						Key:   key,
						Value: val,
					},
				},
			})
		})
	}

	upOpt := options.Update()
	upOpt.SetUpsert(true)
	_, err := db.Table.UpdateMany(db.Ctx, FK, UK, upOpt)
	if err != nil {
		global.LogErr("disposeTask.UpdateEmailCode 数据更插失败", err)
	}
}
