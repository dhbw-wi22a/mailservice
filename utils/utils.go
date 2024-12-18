package utils

import (
	"go_mailservice/config"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.SMTPConfig.Username)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)

	dialer := gomail.NewDialer(
		config.SMTPConfig.Server,
		config.SMTPConfig.Port,
		config.SMTPConfig.Username,
		config.SMTPConfig.Password,
	)

	return dialer.DialAndSend(mailer)
}
