package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"egogo/internal/database"
	"egogo/internal/middleware"
	"egogo/internal/models"
)

func ListTemplates(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var templates []models.Template
	if result := database.DB.Where("user_id = ?", userID).Find(&templates); result.Error != nil {
		ErrorJSON(w, http.StatusInternalServerError, result.Error)
		return
	}

	JSON(w, http.StatusOK, templates)
}

func CreateTemplate(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input struct {
		Name    string `json:"name"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	template := models.Template{
		UserID:       userID,
		Name:         input.Name,
		Subject:      input.Subject,
		Body:         input.Body,
		TriggerToken: uuid.New().String(),
	}

	if result := database.DB.Create(&template); result.Error != nil {
		ErrorJSON(w, http.StatusInternalServerError, result.Error)
		return
	}

	JSON(w, http.StatusCreated, template)
}

func GetTemplate(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid template ID", http.StatusBadRequest)
		return
	}

	var template models.Template
	if result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&template); result.Error != nil {
		ErrorJSON(w, http.StatusNotFound, result.Error)
		return
	}

	JSON(w, http.StatusOK, template)
}

func UpdateTemplate(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid template ID", http.StatusBadRequest)
		return
	}

	var template models.Template
	if result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&template); result.Error != nil {
		ErrorJSON(w, http.StatusNotFound, result.Error)
		return
	}

	var input struct {
		Name    string `json:"name"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	template.Name = input.Name
	template.Subject = input.Subject
	template.Body = input.Body

	if result := database.DB.Save(&template); result.Error != nil {
		ErrorJSON(w, http.StatusInternalServerError, result.Error)
		return
	}

	JSON(w, http.StatusOK, template)
}

func DeleteTemplate(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid template ID", http.StatusBadRequest)
		return
	}

	if result := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Template{}); result.Error != nil {
		ErrorJSON(w, http.StatusInternalServerError, result.Error)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
