package utils

import (
	"encoding/json"
	"errors"
	"go_mailservice/config"
	"gopkg.in/gomail.v2"
)

func ParseRecipients(rawRecipients json.RawMessage) ([]string, error) {
	var recipients []string
	if err := json.Unmarshal(rawRecipients, &recipients); err == nil {
		return recipients, nil
	}

	var singleRecipient string
	if err := json.Unmarshal(rawRecipients, &singleRecipient); err == nil {
		return []string{singleRecipient}, nil
	}
	return nil, errors.New("invalid recipients format: must be a string or array of strings")
}

func SendEmail(to, subject, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.SMTPConfig.Username)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		config.SMTPConfig.Server,
		config.SMTPConfig.Port,
		config.SMTPConfig.Username,
		config.SMTPConfig.Password,
	)

	return dialer.DialAndSend(mailer)
}
