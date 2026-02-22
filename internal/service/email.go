package service

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"sync"
	textTemplate "text/template"

	"egogo/internal/models"
)

type EmailSender interface {
	Send(to, subject, body string) error
	SendFromTemplate(template models.Template, to string, data map[string]interface{}) error
}

// --- Mock Sender ---

type SentEmail struct {
	To      string
	Subject string
	Body    string
}

type MockSender struct {
	mu         sync.Mutex
	SentEmails []SentEmail
}

func (s *MockSender) Send(to, subject, body string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SentEmails = append(s.SentEmails, SentEmail{To: to, Subject: subject, Body: body})
	log.Printf("MOCK EMAIL SENT:\nTo: %s\nSubject: %s\nBody: %s\n", to, subject, body)
	return nil
}

func (s *MockSender) SendFromTemplate(tmpl models.Template, to string, data map[string]interface{}) error {
	// Parse Body (HTML)
	tBody, err := template.New("email_body").Parse(tmpl.Body)
	if err != nil {
		return fmt.Errorf("failed to parse template body: %w", err)
	}
	var body bytes.Buffer
	if err := tBody.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template body: %w", err)
	}

	// Parse Subject (Text)
	tSubject, err := textTemplate.New("email_subject").Parse(tmpl.Subject)
	if err != nil {
		return fmt.Errorf("failed to parse template subject: %w", err)
	}
	var subject bytes.Buffer
	if err := tSubject.Execute(&subject, data); err != nil {
		return fmt.Errorf("failed to execute template subject: %w", err)
	}

	return s.Send(to, subject.String(), body.String())
}

// --- SMTP Sender ---

type SMTPSender struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

func NewSMTPSender(host, port, username, password, from string) *SMTPSender {
	return &SMTPSender{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (s *SMTPSender) Send(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	address := fmt.Sprintf("%s:%s", s.Host, s.Port)

	// Simple email message format
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s", to, subject, body))

	if err := smtp.SendMail(address, auth, s.From, []string{to}, msg); err != nil {
		return fmt.Errorf("failed to send email via SMTP: %w", err)
	}
	return nil
}

func (s *SMTPSender) SendFromTemplate(tmpl models.Template, to string, data map[string]interface{}) error {
	// Parse Body (HTML)
	tBody, err := template.New("email_body").Parse(tmpl.Body)
	if err != nil {
		return fmt.Errorf("failed to parse template body: %w", err)
	}
	var body bytes.Buffer
	if err := tBody.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template body: %w", err)
	}

	// Parse Subject (Text)
	tSubject, err := textTemplate.New("email_subject").Parse(tmpl.Subject)
	if err != nil {
		return fmt.Errorf("failed to parse template subject: %w", err)
	}
	var subject bytes.Buffer
	if err := tSubject.Execute(&subject, data); err != nil {
		return fmt.Errorf("failed to execute template subject: %w", err)
	}

	return s.Send(to, subject.String(), body.String())
}

var Sender EmailSender = &MockSender{} // Default
