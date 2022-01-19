package helper

import (
	"Program/config"
	"Program/constants"
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

func SendEmail(EmailAddress string) ReturnType {
	m := gomail.NewMessage()

	mailConfig := config.GetMailConfig()

	SendTime := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", time.Now().Year(), time.Now().Minute(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	// 获取验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	VerifyCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	html := fmt.Sprintf(`<div>
		<div>
			尊敬的%s, 您好!
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p> 您于 %s 提交了邮箱验证，本次验证码为 %s，为了保证账号安全，验证码有效期为15分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解和使用。 </p>
		</div>
		<div>
			<p> 此邮箱为系统邮箱，请勿回复。</p>
		</div>
	<div>`, EmailAddress, SendTime, VerifyCode)

	m.SetAddressHeader("From", mailConfig["username"].(string), mailConfig["from"].(string))
	m.SetHeader("To", EmailAddress)
	m.SetHeader("Subject", "验证码")
	m.SetBody("text/html", html)
	send := gomail.NewDialer(mailConfig["host"].(string), mailConfig["port"].(int), mailConfig["username"].(string), mailConfig["password"].(string))
	err := send.DialAndSend(m)
	if err != nil {
		return ReturnType{Status: constants.CodeError, Msg: "邮件发送失败", Data: err.Error()}
	} else {
		return ReturnType{Status: constants.CodeSuccess, Msg: "邮件发送成功", Data: ""}
	}
}
