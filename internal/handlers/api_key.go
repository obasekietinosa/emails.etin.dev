package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"egogo/internal/middleware"
	"egogo/internal/models"
)

func (h *Handler) GenerateAPIKey(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input struct {
		Name string `json:"name"`
	}
	// Decode optional input
	_ = json.NewDecoder(r.Body).Decode(&input)
	if input.Name == "" {
		input.Name = "Default API Key"
	}

	keyBytes := make([]byte, 32)
	if _, err := rand.Read(keyBytes); err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	key := hex.EncodeToString(keyBytes)

	apiKey := models.ApiKey{
		UserID: userID,
		Key:    key,
		Name:   input.Name,
	}

	if err := h.Repo.CreateAPIKey(&apiKey); err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusCreated, apiKey)
}

func (h *Handler) ListAPIKeys(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	keys, err := h.Repo.ListAPIKeys(userID)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, keys)
}
