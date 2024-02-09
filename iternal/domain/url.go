package domain

import "github.com/gofrs/uuid"

type Urls struct {
	Base
	UserID uuid.UUID `gorm:"type:uuid;"`
	Url    string    `gorm:"column:url; type:text"`
	Short  string    `gorm:"column:short; type:text"`
}
