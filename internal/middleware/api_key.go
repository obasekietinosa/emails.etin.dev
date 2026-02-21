package middleware

import (
	"context"
	"net/http"

	"github.com/tinrab/emails/internal/database"
	"github.com/tinrab/emails/internal/models"
)

func APIKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			// Check query param as fallback
			apiKey = r.URL.Query().Get("api_key")
		}

		if apiKey == "" {
			http.Error(w, "API Key is required", http.StatusUnauthorized)
			return
		}

		var keyRecord models.ApiKey
		if result := database.DB.Where("key = ?", apiKey).First(&keyRecord); result.Error != nil {
			http.Error(w, "Invalid API Key", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, keyRecord.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
