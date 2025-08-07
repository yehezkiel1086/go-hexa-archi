package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-hexa-archi/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-hexa-archi/internal/core/domain"
	"gorm.io/gorm"
)

func CreateNewRole(c *gin.Context) {
	// connect DB
	db, err := postgres.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return
	}

	// check empty input
	var role domain.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create new role
	if err := db.Create(&role).Error; err != nil {
		if err == gorm.ErrDuplicatedKey {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Role already existed!",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "New role created!",
	})
}
