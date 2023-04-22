package dbType

/*
用来存储 用户信息
db: Account
collection : User
*/

type UserTable struct {
	UserID         string   `bson:"UserID"`         // 用户 ID
	Email          string   `bson:"Email"`          // 用户主要的 Email
	UserEmail      []string `bson:"UserEmail"`      // 用户的 Email 列表
	Avatar         string   `bson:"Avatar"`         // 用户头像
	NickName       string   `bson:"NickName"`       // 用户昵称
	CreateTime     int64    `bson:"CreateTime"`     // 创建时间
	UpdateTime     int64    `bson:"UpdateTime"`     // 更新时间
	EntrapmentCode string   `bson:"EntrapmentCode"` // 防钓鱼码
	Password       string   `bson:"Password"`       // 用户密码
}

/*
用来存储 登录信息
db: Account
collection : LoginInfo
*/
type LoginSucceedType struct {
	UserID         string `bson:"UserID"`
	Email          string `bson:"Email"`
	BrowserName    string `bson:"BrowserName"`
	OsName         string `bson:"OsName"`
	Hostname       string `bson:"Hostname"`
	ISP            string `bson:"ISP"`
	Operators      string `bson:"Operators"`
	CreateTimeUnix int64  `bson:"CreateTimeUnix"`
	CreateTimeStr  string `bson:"CreateTimeStr"`
	Token          string `bson:"Token"`
}
