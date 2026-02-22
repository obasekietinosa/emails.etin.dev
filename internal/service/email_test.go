package service

import (
	"egogo/internal/models"
	"testing"
)

func TestMockSender_SendFromTemplate(t *testing.T) {
	sender := &MockSender{}

	tmpl := models.Template{
		Subject: "Hello {{.Name}}",
		Body:    "Welcome to {{.Service}}!",
	}

	data := map[string]interface{}{
		"Name":    "John",
		"Service": "Egogo",
	}

	err := sender.SendFromTemplate(tmpl, "john@example.com", data)
	if err != nil {
		t.Fatalf("SendFromTemplate failed: %v", err)
	}

	if len(sender.SentEmails) != 1 {
		t.Fatalf("Expected 1 email sent, got %d", len(sender.SentEmails))
	}

	email := sender.SentEmails[0]
	if email.To != "john@example.com" {
		t.Errorf("Expected recipient john@example.com, got %s", email.To)
	}
	if email.Subject != "Hello John" {
		t.Errorf("Expected subject 'Hello John', got '%s'", email.Subject)
	}
	if email.Body != "Welcome to Egogo!" {
		t.Errorf("Expected body 'Welcome to Egogo!', got '%s'", email.Body)
	}
}
