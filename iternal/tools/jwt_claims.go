package tools

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
	Login  string
}
