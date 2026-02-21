package handlers

import (
	"encoding/json"
	"net/http"

	"egogo/internal/database"
	"egogo/internal/middleware"
	"egogo/internal/models"
	"egogo/internal/service"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if input.To == "" {
		http.Error(w, "Recipient 'to' is required", http.StatusBadRequest)
		return
	}

	if err := service.Sender.Send(input.To, input.Subject, input.Body); err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	// Log to database
	logEntry := models.EmailLog{
		UserID:    userID,
		Recipient: input.To,
		Subject:   input.Subject,
		Status:    "sent",
		Mode:      "headless",
	}
	// Best effort logging, don't fail request if logging fails
	go func() {
		database.DB.Create(&logEntry)
	}()

	JSON(w, http.StatusOK, map[string]interface{}{"status": "sent", "user_id": userID})
}
