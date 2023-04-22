package disposeTask

import (
	"Message.net/server/global"
	"Message.net/server/global/dbType"
	"Message.net/server/tmpl"
	"Message.net/server/utils/dbUser"
	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTask"
	"github.com/EasyGolang/goTools/mVerify"
)

func GetUserData(UserID string) (resData dbType.UserTable, resErr error) {
	resData = dbType.UserTable{}
	resErr = nil
	UserDB, err := dbUser.NewUserDB(dbUser.NewUserOpt{
		UserID: UserID,
	})
	if err != nil {
		UserDB.DB.Close()
		resErr = err
		return
	}
	defer UserDB.DB.Close()
	resData = UserDB.Data
	return
}

// ======= 发送系统邮件 =======
func SendSysEmail(TaskCont mTask.SysEmail) error {
	EmailServe := global.GetEmailServe()

	ToEmail := []string{}
	errArr := []error{}
	for _, item := range TaskCont.To {
		isEmail := mVerify.IsEmail(item)
		if isEmail {
			ToEmail = append(ToEmail, item)
			continue
		}
		UserInfo, err := GetUserData(item)
		if err != nil {
			errArr = append(errArr, err)
			continue
		}
		ToEmail = append(ToEmail, UserInfo.Email)
		if len(TaskCont.To) == 1 {
			TaskCont.SendData.EntrapmentCode = UserInfo.EntrapmentCode
		}
	}

	if len(errArr) > 0 {
		global.LogErr("获取邮箱失败:", errArr)
	}

	emailOpt := mEmail.Opt{
		Account:     EmailServe.Account,
		Password:    EmailServe.Password,
		Port:        EmailServe.Port,
		Host:        EmailServe.Host,
		To:          ToEmail,
		From:        TaskCont.From,
		Subject:     TaskCont.Subject,
		TemplateStr: tmpl.SysEmail, //  采用系统模板
		SendData:    TaskCont.SendData,
	}
	// 发送并存储记录
	err := global.SendEmail(emailOpt)
	return err
}

// ======= 发送 验证码 邮件 =======
func SendCodeEmail(TaskCont mTask.CodeEmail) error {
	EmailServe := global.GetEmailServe()
	emailOpt := mEmail.Opt{
		Account:     EmailServe.Account,
		Password:    EmailServe.Password,
		Port:        EmailServe.Port,
		Host:        EmailServe.Host,
		To:          []string{TaskCont.To},
		From:        TaskCont.From,
		Subject:     TaskCont.Subject,
		TemplateStr: tmpl.CodeEmail, //  采用验证码模板
		SendData:    TaskCont.SendData,
	}
	// 发送并存储记录
	err := global.SendEmail(emailOpt)
	// err 为 nil 的时候
	if err == nil {
		if len(TaskCont.SendData.VerifyCode) > 0 && len(TaskCont.To) > 0 {
			UpdateEmailCode(TaskCont)
		} else {
			global.LogErr("disposeTask.EmailAction 空验证码", mJson.Format(TaskCont))
		}
	}

	return err
}

// ======= 发送注册成功邮件 =======
func SendRegisterSucceedEmail(TaskCont mTask.RegisterSucceedEmail) error {
	EmailServe := global.GetEmailServe()
	emailOpt := mEmail.Opt{
		Account:     EmailServe.Account,
		Password:    EmailServe.Password,
		Port:        EmailServe.Port,
		Host:        EmailServe.Host,
		To:          []string{TaskCont.To},
		From:        TaskCont.From,
		Subject:     TaskCont.Subject,
		TemplateStr: tmpl.RegisterSucceedEmail, //  采用系统模板
		SendData:    TaskCont.SendData,
	}
	// 发送并存储记录
	err := global.SendEmail(emailOpt)
	return err
}
