package handlers

import (
	"net/http"

	"egogo/internal/middleware"
)

func (h *Handler) GetEmailLogs(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	logs, err := h.Repo.ListEmailLogs(userID)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, logs)
}
