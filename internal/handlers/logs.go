package handlers

import (
	"net/http"

	"egogo/internal/database"
	"egogo/internal/middleware"
	"egogo/internal/models"
)

func GetEmailLogs(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var logs []models.EmailLog
	// Retrieve logs for the authenticated user, ordered by creation time descending
	if result := database.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&logs); result.Error != nil {
		ErrorJSON(w, http.StatusInternalServerError, result.Error)
		return
	}

	JSON(w, http.StatusOK, logs)
}
