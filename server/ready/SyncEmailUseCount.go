package ready

import (
	"Message.net/server/global"
	"Message.net/server/global/config"
	"Message.net/server/global/dbType"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mTime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 同步账号的发送频率
func SyncEmailUseCount() {
	db := mMongo.New(mMongo.Opt{
		UserName: config.SysEnv.MongoUserName,
		Password: config.SysEnv.MongoPassword,
		Address:  config.SysEnv.MongoAddress,
		DBName:   "Message",
	}).Connect().Collection("Email")
	defer global.Run.Println("global.SyncEmailUseCount 同步一次发信频率")
	defer db.Close()

	UseEmailCountHour := make(map[string]int)
	UseEmailCount24Hour := make(map[string]int)

	global.EmailCount = make(map[string]global.EmailCountType)

	HourList := HourQuery(db, 1)
	Hour24List := HourQuery(db, 24)

	for _, val := range HourList {
		nowCount := UseEmailCountHour[val.Account]
		nowCount++
		UseEmailCountHour[val.Account] = nowCount
	}

	for _, val := range Hour24List {
		nowCount := UseEmailCount24Hour[val.Account]
		nowCount++
		UseEmailCount24Hour[val.Account] = nowCount
	}

	for _, val := range global.EmailAccountList {
		global.EmailCount[val.Account] = global.EmailCountType{
			Hour:   UseEmailCountHour[val.Account],
			Hour24: UseEmailCount24Hour[val.Account],
		}
	}

	global.Run.Println("同步发信频率",
		mJson.Format(global.EmailCount),
	)
}

func HourQuery(db *mMongo.DB, hour int64) []dbType.MessageEmail {
	findOpt := options.Find()
	findOpt.SetSort(map[string]int{
		"CreateTime": -1,
	})
	findOpt.SetAllowDiskUse(true)

	now := mTime.GetUnixInt64()                                         // 当前时间
	hourAgo := mTime.GetUnixInt64() - (hour * mTime.UnixTimeInt64.Hour) //  多少小时之前

	// 最近 1小时 的查询
	FK := bson.D{{
		Key: "CreateTime",
		Value: bson.D{
			{
				Key:   "$gte", // 大于或等于
				Value: hourAgo,
			}, {
				Key:   "$lte", // 小于或等于
				Value: now,
			},
		},
	}}

	cur, err := db.Table.Find(db.Ctx, FK, findOpt)
	if err != nil {
		db.Close()
		global.LogErr("SyncEmailUseCount,数据库读取失败 %+v", err)
		return nil
	}
	// 1 小时的数据更新
	var EmailAccountList []dbType.MessageEmail
	for cur.Next(db.Ctx) {
		var curData dbType.MessageEmail
		cur.Decode(&curData)

		EmailAccountList = append(EmailAccountList, curData)
	}
	return EmailAccountList
}
