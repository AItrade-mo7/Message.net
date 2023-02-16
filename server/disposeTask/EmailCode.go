package disposeTask

import (
	"fmt"

	"Message.net/server/global"
	"Message.net/server/global/config"
	"github.com/EasyGolang/goTools/mMongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateEmailCode(to []string, code string) {
	fmt.Println("存储验证码", to, code)
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
	for _, val := range to {

		FK = append(FK, bson.E{
			Key:   val,
			Value: code,
		})
		UK = append(UK, bson.E{
			Key: "$set",
			Value: bson.D{
				{
					Key:   val,
					Value: code,
				},
			},
		})
	}

	upOpt := options.Update()
	upOpt.SetUpsert(true)
	_, err := db.Table.UpdateMany(db.Ctx, FK, UK, upOpt)
	if err != nil {
		global.LogErr("disposeTask.UpdateEmailCode 数据更插失败", err)
	}
}
