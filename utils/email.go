package utils

import (
	"github.com/go-gomail/gomail"
	"main/config"
	"strings"
)

type EmailParam struct {
	// ServerHost 邮箱服务器地址，如bridge邮箱为smtp-mail.outlook.com
	ServerHost string
	// ServerPort 邮箱服务器端口，如bridge邮箱为587
	ServerPort int
	// FromEmail　发件人邮箱地址
	FromEmail string
	// FromPasswd 发件人邮箱密码（注意，这里是明文形式），TODO：如果设置成密文？
	FromPasswd string
	// Toers 接收者邮件，如有多个，则以英文逗号(“,”)隔开，不能为空
	Toers string
	// CCers 抄送者邮件，如有多个，则以英文逗号(“,”)隔开，可以为空
	CCers string
	//用户识别CD(uuid)
	ActiveCode string
	//邮件主体
	Message *gomail.Message
}

func InitEmail(ep *EmailParam) {
	toers := []string{}

	ep.Message = gomail.NewMessage()

	if len(ep.Toers) == 0 {
		return
	}

	for _, tmp := range strings.Split(ep.Toers, ",") {
		toers = append(toers, strings.TrimSpace(tmp))
	}

	// 收件人可以有多个，故用此方式
	ep.Message.SetHeader("To", toers...)

	//抄送列表
	if len(ep.CCers) != 0 {
		for _, tmp := range strings.Split(ep.CCers, ",") {
			toers = append(toers, strings.TrimSpace(tmp))
		}
		ep.Message.SetHeader("Cc", toers...)
	}

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	ep.Message.SetAddressHeader("From", ep.FromEmail, "")

	// 主题
	ep.Message.SetHeader("Subject", "From Bridge System")
	//str := fmt.Formatter()
	// 正文
	message :=
		"<p>ご本人様確認のため、<a href=\"http://" + config.InitConfig().Host + "/#/Active" + "?active_code=" + ep.ActiveCode + "\">ここ</a>をクリックし</p>" +
			"<p>アカウントの認証を完了させて下さい。</p>" +
			"By BRS セキュリティサービス</p>"

	ep.Message.SetBody("text/html", message)

	dialer := gomail.NewPlainDialer(ep.ServerHost, ep.ServerPort, ep.FromEmail, ep.FromPasswd)
	// 发送
	err := dialer.DialAndSend(ep.Message)
	if err != nil {
		panic(err)
	}

}
