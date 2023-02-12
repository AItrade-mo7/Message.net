package dbType

import "github.com/EasyGolang/goTools/mEmail"

// ======= Message 数据库 ==========
// Email
type MessageEmail struct {
	mEmail.Opt
	EmailID    string `bson:"EmailID"`
	SendResult string `bson:"SendResult"`
}
