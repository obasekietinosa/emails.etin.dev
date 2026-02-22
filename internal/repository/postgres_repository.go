package repository

import (
	"egogo/internal/models"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	DB *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

// User methods
func (r *PostgresRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *PostgresRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// API Key methods
func (r *PostgresRepository) CreateAPIKey(apiKey *models.ApiKey) error {
	return r.DB.Create(apiKey).Error
}

func (r *PostgresRepository) ListAPIKeys(userID uint) ([]models.ApiKey, error) {
	var keys []models.ApiKey
	if err := r.DB.Where("user_id = ?", userID).Find(&keys).Error; err != nil {
		return nil, err
	}
	return keys, nil
}

func (r *PostgresRepository) GetAPIKeyByValue(key string) (*models.ApiKey, error) {
	var apiKey models.ApiKey
	if err := r.DB.Where("key = ?", key).First(&apiKey).Error; err != nil {
		return nil, err
	}
	return &apiKey, nil
}

// Template methods
func (r *PostgresRepository) CreateTemplate(template *models.Template) error {
	return r.DB.Create(template).Error
}

func (r *PostgresRepository) ListTemplates(userID uint) ([]models.Template, error) {
	var templates []models.Template
	if err := r.DB.Where("user_id = ?", userID).Find(&templates).Error; err != nil {
		return nil, err
	}
	return templates, nil
}

func (r *PostgresRepository) GetTemplate(id uint, userID uint) (*models.Template, error) {
	var template models.Template
	if err := r.DB.Where("id = ? AND user_id = ?", id, userID).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

func (r *PostgresRepository) UpdateTemplate(template *models.Template) error {
	return r.DB.Save(template).Error
}

func (r *PostgresRepository) DeleteTemplate(id uint, userID uint) error {
	return r.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Template{}).Error
}

func (r *PostgresRepository) GetTemplateByToken(token string) (*models.Template, error) {
	var template models.Template
	if err := r.DB.Where("trigger_token = ?", token).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

// Email Log methods
func (r *PostgresRepository) CreateEmailLog(log *models.EmailLog) error {
	return r.DB.Create(log).Error
}

func (r *PostgresRepository) ListEmailLogs(userID uint) ([]models.EmailLog, error) {
	var logs []models.EmailLog
	if err := r.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
