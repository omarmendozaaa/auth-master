package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Description  string    `json:"description"`
	ClientID     string    `gorm:"uniqueIndex;not null" json:"client_id"`
	ClientSecret string    `gorm:"uniqueIndex;not null" json:"client_secret"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (project *Project) BeforeCreate(tx *gorm.DB) (err error) {
	project.ClientID = uuid.New().String()
	project.ClientSecret = uuid.New().String()
	return
}
