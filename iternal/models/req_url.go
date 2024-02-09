package models

import "github.com/gofrs/uuid"

type ReqUrl struct {
	UserID uuid.UUID `json:"user_id,omitempty"`
	Url    string    `json:"url" binding:"required"`
}
