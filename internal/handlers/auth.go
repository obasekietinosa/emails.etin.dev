package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/my-org/emails-service/internal/database"
	"github.com/my-org/emails-service/internal/models"
	"github.com/my-org/emails-service/internal/service"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct{}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
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

	if result := database.DB.Create(&user); result.Error != nil {
		ErrorJSON(w, http.StatusInternalServerError, result.Error)
		return
	}

	token, err := service.GenerateToken(&user)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusCreated, map[string]string{"token": token})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	var user models.User
	if result := database.DB.Where("email = ?", input.Email).First(&user); result.Error != nil {
		ErrorJSON(w, http.StatusUnauthorized, result.Error)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	token, err := service.GenerateToken(&user)
	if err != nil {
		ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, map[string]string{"token": token})
}
