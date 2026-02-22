package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"egogo/internal/middleware"
	"egogo/internal/models"
	"egogo/internal/repository"
	"egogo/internal/service"
)

func TestCreateTemplate(t *testing.T) {
	// Setup
	repo := repository.NewMockRepository()
	sender := &service.MockSender{}
	h := NewHandler(repo, sender)

	// Prepare request
	input := map[string]string{
		"name":    "Welcome Email",
		"subject": "Welcome {{.Name}}",
		"body":    "Hello {{.Name}}",
	}
	body, _ := json.Marshal(input)
	req := httptest.NewRequest("POST", "/templates", bytes.NewBuffer(body))

	// Add UserID to context
	ctx := context.WithValue(req.Context(), middleware.UserContextKey, uint(1))
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	// Execute
	h.CreateTemplate(rr, req)

	// Assert
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status 201 Created, got %d", rr.Code)
	}

	var response models.Template
	json.NewDecoder(rr.Body).Decode(&response)

	if response.Name != input["name"] {
		t.Errorf("Expected name %s, got %s", input["name"], response.Name)
	}
	if response.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", response.UserID)
	}
	if response.TriggerToken == "" {
		t.Error("Expected TriggerToken to be generated")
	}

	// Verify it's in the repo
	tmpl, err := repo.GetTemplate(response.ID, 1)
	if err != nil {
		t.Fatalf("Template not found in repo: %v", err)
	}
	if tmpl.Name != input["name"] {
		t.Errorf("Repo template name mismatch")
	}
}
