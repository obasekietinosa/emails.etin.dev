package repository

import (
	"errors"
	"egogo/internal/models"
)

var ErrRecordNotFound = errors.New("record not found")

type MockRepository struct {
	Users     map[string]*models.User
	Templates map[uint]*models.Template
	APIKeys   map[string]*models.ApiKey
	EmailLogs []models.EmailLog
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		Users:     make(map[string]*models.User),
		Templates: make(map[uint]*models.Template),
		APIKeys:   make(map[string]*models.ApiKey),
		EmailLogs: []models.EmailLog{},
	}
}

func (m *MockRepository) CreateUser(user *models.User) error {
	if m.Users == nil {
		m.Users = make(map[string]*models.User)
	}
	m.Users[user.Email] = user
	return nil
}

func (m *MockRepository) GetUserByEmail(email string) (*models.User, error) {
	if user, ok := m.Users[email]; ok {
		return user, nil
	}
	return nil, ErrRecordNotFound
}

func (m *MockRepository) GetUserByID(id uint) (*models.User, error) {
	// Not implemented for now unless needed
	return nil, ErrRecordNotFound
}

func (m *MockRepository) CreateAPIKey(apiKey *models.ApiKey) error {
	if m.APIKeys == nil {
		m.APIKeys = make(map[string]*models.ApiKey)
	}
	m.APIKeys[apiKey.Key] = apiKey
	return nil
}

func (m *MockRepository) ListAPIKeys(userID uint) ([]models.ApiKey, error) {
	var keys []models.ApiKey
	for _, k := range m.APIKeys {
		if k.UserID == userID {
			keys = append(keys, *k)
		}
	}
	return keys, nil
}

func (m *MockRepository) GetAPIKeyByValue(key string) (*models.ApiKey, error) {
	if k, ok := m.APIKeys[key]; ok {
		return k, nil
	}
	return nil, ErrRecordNotFound
}

func (m *MockRepository) CreateTemplate(template *models.Template) error {
	if m.Templates == nil {
		m.Templates = make(map[uint]*models.Template)
	}
	// Simulate ID generation
	if template.ID == 0 {
		template.ID = uint(len(m.Templates) + 1)
	}
	m.Templates[template.ID] = template
	return nil
}

func (m *MockRepository) ListTemplates(userID uint) ([]models.Template, error) {
	var templates []models.Template
	for _, t := range m.Templates {
		if t.UserID == userID {
			templates = append(templates, *t)
		}
	}
	return templates, nil
}

func (m *MockRepository) GetTemplate(id uint, userID uint) (*models.Template, error) {
	if t, ok := m.Templates[id]; ok && t.UserID == userID {
		return t, nil
	}
	return nil, ErrRecordNotFound
}

func (m *MockRepository) UpdateTemplate(template *models.Template) error {
	m.Templates[template.ID] = template
	return nil
}

func (m *MockRepository) DeleteTemplate(id uint, userID uint) error {
	delete(m.Templates, id)
	return nil
}

func (m *MockRepository) GetTemplateByToken(token string) (*models.Template, error) {
	for _, t := range m.Templates {
		if t.TriggerToken == token {
			return t, nil
		}
	}
	return nil, ErrRecordNotFound
}

func (m *MockRepository) CreateEmailLog(log *models.EmailLog) error {
	m.EmailLogs = append(m.EmailLogs, *log)
	return nil
}

func (m *MockRepository) ListEmailLogs(userID uint) ([]models.EmailLog, error) {
	var logs []models.EmailLog
	for _, l := range m.EmailLogs {
		if l.UserID == userID {
			logs = append(logs, l)
		}
	}
	return logs, nil
}
