package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"egogo/internal/database"
	"egogo/internal/models"
	"egogo/internal/service"
)

func TriggerEmail(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	var template models.Template
	if result := database.DB.Where("trigger_token = ?", token).First(&template); result.Error != nil {
		ErrorJSON(w, http.StatusNotFound, result.Error)
		return
	}

	var input map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		// If body is empty or invalid JSON, we might just proceed with empty map
		// But usually we expect at least empty JSON object "{}" or some fields.
		// Let's assume input is optional or can be empty.
		input = make(map[string]interface{})
	}

	// We need a recipient email address.
	// The prompt says: "In fully managed mode, the user provides a template they want sent to a specific email address or group of email addresses when a trigger is hit"
	// This implies the recipient is either fixed in the template or passed in the trigger request?
	// "when a trigger is hit and can specify variables to be interpolated into this template."
	// Usually, the recipient is also a variable or parameter.
	// Let's assume the recipient is passed in the body as `to` field, or we can look it up if we had a contact list.
	// But `models.Template` doesn't have a `To` field.
	// If the user provides the template, maybe they also provide the default recipient?
	// Or maybe the recipient is passed in the trigger payload.
	// Let's assume `to` is in the input map.

	to, ok := input["to"].(string)
	if !ok || to == "" {
		http.Error(w, "Recipient 'to' is required in request body", http.StatusBadRequest)
		return
	}

	if err := service.Sender.SendFromTemplate(template, to, input); err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	// Log to database
	logEntry := models.EmailLog{
		UserID:    template.UserID,
		Recipient: to,
		Subject:   template.Subject,
		Status:    "sent",
		Mode:      "managed",
	}
	go func() {
		database.DB.Create(&logEntry)
	}()

	JSON(w, http.StatusOK, map[string]string{"status": "sent"})
}
