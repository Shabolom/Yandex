package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type SaveStudent struct {
	ID        uuid.UUID `gorm:"type:uuid;" json:"id"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
	Name      string    `json:"name,omitempty" binding:"required"`
	Surname   string    `json:"surname" binding:"required"`
	Gender    string    `json:"gender" binding:"required"`
	Age       int       `json:"age" binding:"required"`
	Country   string    `json:"country" binding:"required"`
	Email     string    `json:"email" binding:"required"`
}
