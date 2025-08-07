package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yehezkiel1086/go-hexa-archi/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-hexa-archi/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// connect DB
	db, err := postgres.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return
	}

	// check empty user input
	user := &domain.User{
		ID: uuid.New(),
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username, password, name, email and phone are required",
		})
		return
	}

	// encrypt password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password encryption failed",
		})
		return
	}

	user.Password = string(hashedPwd)

	// use user as default role
	var userRole domain.Role

	// default role is user
	if err := db.Model(&domain.Role{}).Where("role = ?", "user").First(&userRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.RoleID = userRole.ID

	// create new user
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}
