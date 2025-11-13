package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProjectID   uuid.UUID
	Name        string `gorm:"not null"`
	Description string
	IsDefault   bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
