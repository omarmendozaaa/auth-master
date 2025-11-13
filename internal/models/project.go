package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string    `gorm:"not null"`
	Description  string
	ClientID     string    `gorm:"uniqueIndex;not null"`
	ClientSecret string    `gorm:"uniqueIndex;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func (project *Project) BeforeCreate(tx *gorm.DB) (err error) {
	project.ClientID = uuid.New().String()
	project.ClientSecret = uuid.New().String()
	return
}
