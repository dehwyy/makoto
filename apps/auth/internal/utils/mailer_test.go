package utils

import (
	"fmt"
	"testing"

	mailtemplates "github.com/dehwyy/makoto/apps/auth/internal/utils/mail-templates"
	"github.com/dehwyy/makoto/libs/config"
)

func TestRequest(t *testing.T) {
	cfg := config.New()
	sender := NewGmailSender(cfg.GmailSennderName, cfg.GmailSenderAddr, cfg.GmailSenderPassword)

	subject := "TEST:Makoto. Confirm your email"
	html, err := mailtemplates.ParseTemplate(mailtemplates.Confirm, &mailtemplates.ConfirmData{
		Username: "dehwyy",
		Link:     "https://google.com",
	})
	if err != nil {
		fmt.Printf("Error template %v\n", err)
		return
	}

	to := "dehwyy@yandex.ru"

	err = sender.Send(subject, html, to)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	} else {
		fmt.Printf("Success\n")
	}
}
