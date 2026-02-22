package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"egogo/internal/models"
)

func (h *Handler) TriggerEmail(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	template, err := h.Repo.GetTemplateByToken(token)
	if err != nil {
		ErrorJSON(w, http.StatusNotFound, err)
		return
	}

	var input map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		input = make(map[string]interface{})
	}

	to, ok := input["to"].(string)
	if !ok || to == "" {
		http.Error(w, "Recipient 'to' is required in request body", http.StatusBadRequest)
		return
	}

	if err := h.Sender.SendFromTemplate(*template, to, input); err != nil {
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
		h.Repo.CreateEmailLog(&logEntry)
	}()

	JSON(w, http.StatusOK, map[string]string{"status": "sent"})
}
