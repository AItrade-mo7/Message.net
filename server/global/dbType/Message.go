package dbType

// ======= Message 数据库 ==========
// Email
type MessageEmail struct {
	// 自定义
	SendResult string `bson:"SendResult"`
	EmailID    string `bson:"EmailID"`
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
