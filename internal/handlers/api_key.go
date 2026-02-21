package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/my-org/emails-service/internal/database"
	"github.com/my-org/emails-service/internal/middleware"
	"github.com/my-org/emails-service/internal/models"
)

func GenerateAPIKey(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
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
		Name:   "Default API Key", // Could be parameterized
	}

	if result := database.DB.Create(&apiKey); result.Error != nil {
		ErrorJSON(w, http.StatusInternalServerError, result.Error)
		return
	}

	JSON(w, http.StatusCreated, apiKey)
}
