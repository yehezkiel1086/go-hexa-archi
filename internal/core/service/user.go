package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yehezkiel1086/go-hexa-archi/internal/adapter/storage/postgres"
)

type UserOutput struct {
	ID        uuid.UUID `json:"id" binding:"required" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string `json:"username" binding:"required"`
	Role string `json:"role" binding:"required" gorm:"size:255;not null;unique"`

	Fullname string `json:"fullname" binding:"required" gorm:"size:255;not null"`
	Email string `json:"email" binding:"required" gorm:"size:255;not null"`
	Phone string `json:"phone" binding:"required" gorm:"size:255;not null"`
}

func GetAllUsers(c *gin.Context) {
	// connect DB
	db, err := postgres.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return
	}

	// get all users
	var users []UserOutput

	rows, err := db.Table("users").Select("users.id, users.username, roles.role, users.fullname, users.email, users.phone").Joins("left join roles on roles.id = users.role_id").Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		db.ScanRows(rows, &users)
	}

	c.JSON(http.StatusOK, users)
}

func GetUserByUsername(c *gin.Context) {
	// connect DB
	db, err := postgres.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return
	}

	// get username from param
	userParam := c.Param("user")
	if userParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User param is required",
		})
		return
	}

	// get user
	var user UserOutput

	if err := db.Table("users").Select("users.id, users.username, roles.role, users.fullname, users.email, users.phone").Joins("left join roles on roles.id = users.role_id").Where("users.username = ?", userParam).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
