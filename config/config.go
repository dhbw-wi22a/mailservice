package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type smtpConfig struct {
	Server      string
	Port        int
	Username    string
	Password    string
	SenderEmail string
}

var SMTPConfig smtpConfig

func init() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	SMTPConfig = smtpConfig{
		Server:      os.Getenv("SMTP_SERVER"),
		Port:        port,
		Username:    os.Getenv("SMTP_USERNAME"),
		Password:    os.Getenv("SMTP_PASSWORD"),
		SenderEmail: os.Getenv("SMTP_SENDER_EMAIL"),
	}

}
