package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	Templates []Template     `gorm:"foreignKey:UserID" json:"templates,omitempty"`
	ApiKeys   []ApiKey       `gorm:"foreignKey:UserID" json:"api_keys,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Template struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	UserID       uint           `gorm:"not null" json:"user_id"`
	Name         string         `gorm:"not null" json:"name"`
	Subject      string         `gorm:"not null" json:"subject"`
	Body         string         `gorm:"not null" json:"body"`
	TriggerToken string         `gorm:"uniqueIndex;not null" json:"trigger_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type ApiKey struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	Key       string         `gorm:"uniqueIndex;not null" json:"key"`
	Name      string         `gorm:"not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type EmailLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Recipient string    `gorm:"not null" json:"recipient"`
	Subject   string    `gorm:"not null" json:"subject"`
	Status    string    `gorm:"not null" json:"status"` // e.g., "sent", "failed"
	Mode      string    `gorm:"not null" json:"mode"`   // "headless" or "managed"
	CreatedAt time.Time `json:"created_at"`
}
