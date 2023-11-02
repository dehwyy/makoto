package mailtemplates

import (
	"bytes"
	"html/template"
)

type ConfirmData struct {
	Username string
	Link     string
}

type MailTemplates string

const (
	Confirm MailTemplates = "confirm.html"
)

func ParseTemplate(template_file MailTemplates, data *ConfirmData) ([]byte, error) {
	t, err := template.ParseFiles("mail-templates/" + string(template_file))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
