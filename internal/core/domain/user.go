package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `json:"id" binding:"required" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Username string `json:"username" binding:"required" gorm:"size:255;not null;unique"`
	Password string `json:"password" binding:"required" gorm:"size:255;not null"`

	Fullname string `json:"fullname" binding:"required" gorm:"size:255;not null"`
	Email string `json:"email" binding:"required" gorm:"size:255;not null"`
	Phone string `json:"phone" binding:"required" gorm:"size:255;not null"`

	RoleID uuid.UUID `json:"role_id" gorm:"type:uuid;not null"`
	UserRole Role	`json:"-" binding:"-" gorm:"foreignKey:RoleID"`
}
