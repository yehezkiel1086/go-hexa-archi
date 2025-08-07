package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	Username string    `json:"username"`
	RoleID   uuid.UUID `json:"role_id" binding:"required" gorm:"type:uuid"`
	jwt.RegisteredClaims
}
