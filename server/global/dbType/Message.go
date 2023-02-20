package dbType

/*
用来存储邮件发送记录
db: Message
collection : Email
*/
type MessageEmail struct {
	// 自定义
	SendResult    string `bson:"SendResult"`
	EmailID       string `bson:"EmailID"`
	CreateTime    int64  `bson:"CreateTime"`
	CreateTimeStr string `bson:"CreateTimeStr"`
	// 来自 mEmail.Opt
	Account     string   `bson:"Account"`
	Password    string   `bson:"Password"`
	To          []string `bson:"To"`
	From        string   `bson:"From"`
	Subject     string   `bson:"Subject"`
	Port        string   `bson:"Port"`
	Host        string   `bson:"Host"`
	TemplateStr string   `bson:"TemplateStr"`
	SendData    any      `bson:"SendData"`
}
