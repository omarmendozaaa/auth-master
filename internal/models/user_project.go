package models

import (
	"time"

	"github.com/google/uuid"
)

type UserProject struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID
	ProjectID uuid.UUID
	RoleID    uuid.UUID
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
