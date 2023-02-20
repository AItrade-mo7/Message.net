package dbType

/*
用来存储验证码
db: Message
collection : VerifyCode
*/
type EmailCodeTable struct {
	Email    string `bson:"Email"`
	Code     string `bson:"Code"`
	SendTime int64  `bson:"SendTime"`
}
