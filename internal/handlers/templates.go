package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"egogo/internal/middleware"
	"egogo/internal/models"
)

func (h *Handler) ListTemplates(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey).(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	templates, err := h.Repo.ListTemplates(userID)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, templates)
}

func (h *Handler) CreateTemplate(w http.ResponseWriter, r *http.Request) {
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

	if err := h.Repo.CreateTemplate(&template); err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusCreated, template)
}

func (h *Handler) GetTemplate(w http.ResponseWriter, r *http.Request) {
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

	template, err := h.Repo.GetTemplate(uint(id), userID)
	if err != nil {
		ErrorJSON(w, http.StatusNotFound, err)
		return
	}

	JSON(w, http.StatusOK, template)
}

func (h *Handler) UpdateTemplate(w http.ResponseWriter, r *http.Request) {
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

	template, err := h.Repo.GetTemplate(uint(id), userID)
	if err != nil {
		ErrorJSON(w, http.StatusNotFound, err)
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

	if err := h.Repo.UpdateTemplate(template); err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, template)
}

func (h *Handler) DeleteTemplate(w http.ResponseWriter, r *http.Request) {
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

	if err := h.Repo.DeleteTemplate(uint(id), userID); err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
