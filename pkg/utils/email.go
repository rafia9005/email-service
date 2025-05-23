package utils

import (
    "fmt"
    "net/smtp"
    "os"
)

func EmailSender(to string, subject string, body string, isHTML bool) error {
    host := os.Getenv("SMTP_HOST")
    port := os.Getenv("SMTP_PORT")
    senderName := os.Getenv("SMTP_SENDER_NAME")
    authEmail := os.Getenv("SMTP_EMAIL")
    authPassword := os.Getenv("SMTP_PASSWORD")

    addr := host + ":" + port
    auth := smtp.PlainAuth("", authEmail, authPassword, host)

    contentType := "text/plain; charset=\"utf-8\""
    if isHTML {
        contentType = "text/html; charset=\"utf-8\""
    }

    msg := "From: " + senderName + " <" + authEmail + ">\r\n" +
        "To: " + to + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "Content-Type: " + contentType + "\r\n" +
        "\r\n" + body + "\r\n"

    err := smtp.SendMail(addr, auth, authEmail, []string{to}, []byte(msg))
    if err != nil {
        return fmt.Errorf("failed to send email: %w", err)
    }

    return nil
}

