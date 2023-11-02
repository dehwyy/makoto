package utils

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailSender interface {
	Send(subject string, html []byte, to string) error
}

type gmail_sender struct {
	from_name     string
	from_addr     string
	from_password string
}

const (
	smtp_auth_gmail_addr   = "smtp.gmail.com"
	smtp_server_gmail_addr = "smtp.gmail.com:587"
)

func NewGmailSender(name, addr, password string) EmailSender {
	return &gmail_sender{
		from_name:     name,
		from_addr:     addr,
		from_password: password,
	}
}

func (s *gmail_sender) Send(subject string, html []byte, to string) error {
	mail := email.NewEmail()

	mail.From = fmt.Sprintf("%s <%s>", s.from_name, s.from_addr)

	mail.Subject = subject
	mail.To = []string{to}
	mail.HTML = html

	smpt_auth := smtp.PlainAuth("", s.from_addr, s.from_password, smtp_auth_gmail_addr)
	return mail.Send(smtp_server_gmail_addr, smpt_auth)
}
