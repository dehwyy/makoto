package utils

import (
	"fmt"
	"os"
	"testing"

	mailtemplates "github.com/dehwyy/makoto/apps/auth/internal/utils/mail-templates"
	"github.com/dehwyy/makoto/libs/config"
)

func TestRequest(t *testing.T) {
	// if test is running on github workflow -> skip
	if is_github_workflow := os.Getenv("WORKFLOW"); is_github_workflow != "" {
		return
	}

	fmt.Println("Sending email...")

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

	// well its mine address
	to := "dehwyy@yandex.ru"

	err = sender.Send(subject, html, to)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	} else {
		fmt.Printf("Success!\n")
	}

	fmt.Println("Email sent!")
}
