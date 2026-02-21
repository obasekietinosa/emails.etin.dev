package service

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"egogo/internal/models"
)

type EmailSender interface {
	Send(to, subject, body string) error
	SendFromTemplate(template models.Template, to string, data map[string]interface{}) error
}

type MockSender struct{}

func (s *MockSender) Send(to, subject, body string) error {
	log.Printf("MOCK EMAIL SENT:\nTo: %s\nSubject: %s\nBody: %s\n", to, subject, body)
	return nil
}

func (s *MockSender) SendFromTemplate(tmpl models.Template, to string, data map[string]interface{}) error {
	t, err := template.New("email").Parse(tmpl.Body)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return s.Send(to, tmpl.Subject, body.String())
}

var Sender EmailSender = &MockSender{}
