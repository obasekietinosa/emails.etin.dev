package repository

import (
	"egogo/internal/models"
)

type Repository interface {
	// User methods
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)

	// API Key methods
	CreateAPIKey(apiKey *models.ApiKey) error
	ListAPIKeys(userID uint) ([]models.ApiKey, error)
	GetAPIKeyByValue(key string) (*models.ApiKey, error) // Need this for headless send? Let's check send handler.

	// Template methods
	CreateTemplate(template *models.Template) error
	ListTemplates(userID uint) ([]models.Template, error)
	GetTemplate(id uint, userID uint) (*models.Template, error)
	UpdateTemplate(template *models.Template) error
	DeleteTemplate(id uint, userID uint) error
	GetTemplateByToken(token string) (*models.Template, error)

	// Email Log methods
	CreateEmailLog(log *models.EmailLog) error
	ListEmailLogs(userID uint) ([]models.EmailLog, error)
}
