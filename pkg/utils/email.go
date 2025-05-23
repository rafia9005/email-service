package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

var (
	CONFIG_SMTP_HOST          = os.Getenv("SMTP_HOST")
	CONFIG_SMTP_PORT          = os.Getenv("SMTP_PORT")
	CONFIG_SMTP_SENDER_NAME   = os.Getenv("SMTP_SENDER_NAME")
	CONFIG_SMTP_AUTH_EMAIL    = os.Getenv("SMTP_EMAIL")
	CONFIG_SMTP_AUTH_PASSWORD = os.Getenv("SMTP_PASSWORD")
)

func EmailSender(to string, subject string, body string, isHTML bool) error {
	addr := CONFIG_SMTP_HOST + ":" + CONFIG_SMTP_PORT

	auth := smtp.PlainAuth("", CONFIG_SMTP_AUTH_EMAIL, CONFIG_SMTP_AUTH_PASSWORD, CONFIG_SMTP_HOST)

	contentType := "text/plain; charset=\"utf-8\""
	if isHTML {
		contentType = "text/html; charset=\"utf-8\""
	}

	msg := "From: " + CONFIG_SMTP_SENDER_NAME + " <" + CONFIG_SMTP_AUTH_EMAIL + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: " + contentType + "\r\n" +
		"\r\n" +
		body + "\r\n"

	err := smtp.SendMail(addr, auth, CONFIG_SMTP_AUTH_EMAIL, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
