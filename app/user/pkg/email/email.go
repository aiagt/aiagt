package email

import (
	"fmt"
	"github.com/aiagt/aiagt/app/user/conf"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func Send(subject, html string, toEmails ...string) error {
	config := conf.Conf().Email
	e := email.NewEmail()

	e.From = config.EmailFrom
	e.To = toEmails
	e.Subject = subject
	e.HTML = []byte(html)
	auth := smtp.PlainAuth("", config.EmailAddress, config.Auth, config.SmtpHost)

	return e.Send(config.SmtpAddr, auth)
}

func SendAuthCaptcha(captcha string, toEmails ...string) error {
	subject := "【Aiagt】邮箱验证"
	html := fmt.Sprintf(`<div style="text-align: center;">
		<h2 style="color: #333;">欢迎使用，你的验证码为：</h2>
		<h1 style="margin: 1.2em 0;">%s</h1>
		<p style="font-size: 12px; color: #666;">请在5分钟内完成验证，过期失效，请勿告知他人，以防个人信息泄露</p>
	</div>`, captcha)
	return Send(subject, html, toEmails...)
}

func SendResetCaptcha(captcha string, toEmails ...string) error {
	subject := "【Aiagt】重置密码"
	html := fmt.Sprintf(`<div style="text-align: center;">
		<h3>您好，<br>您正在重置与此电子邮件地址关联的 Aiagt 帐户的密码，以下是您的验证码：</h3>
		<h1 style="margin: 1.2em 0;">%s</h1>
		<p style="font-size: 12px; color: #666;">请在5分钟内完成验证，过期失效，请勿告知他人，以防个人信息泄露</p>
	</div>`, captcha)
	return Send(subject, html, toEmails...)
}
