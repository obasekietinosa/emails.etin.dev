package handlers

import (
	"encoding/json"
	"net/http"

	"egogo/internal/models"
	"egogo/internal/service"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := h.Repo.CreateUser(&user); err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	token, err := service.GenerateToken(&user)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusCreated, map[string]string{"token": token})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.Repo.GetUserByEmail(input.Email)
	if err != nil {
		// Could be record not found or other error
		ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	token, err := service.GenerateToken(user)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, map[string]string{"token": token})
}
