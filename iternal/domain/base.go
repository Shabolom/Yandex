package domain

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID      `gorm:"type:uuid;" gorm:"primaryKey" json:"id" gorm:"index:id"`
	CreatedAt time.Time      `json:"created-at"`
	UpdatedAt time.Time      `json:"updated-at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
