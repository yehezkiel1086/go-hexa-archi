package service

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yehezkiel1086/go-hexa-archi/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-hexa-archi/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	// connect DB
	db, err := postgres.ConnDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "DB Connection failed",
		})
		return
	}

	// check empty user input
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username and password are required",
		})
		return
	}

	// check username
	var user domain.User

	if err := db.First(&user, "username = ?", input.Username).Error; err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials",
		})		
		return
	}

	// get important things from .env
	strTokenDuration := os.Getenv("TOKEN_DURATION")
	jwtSecret := os.Getenv("JWT_SECRET")
	httpHost := os.Getenv("HTTP_HOST")
	httpPort := os.Getenv("HTTP_PORT")

	if jwtSecret == "" || strTokenDuration == "" || httpHost == "" || httpPort == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "JWT secret and token duration is required from .env",
		})
		return
	}

	// make http_url
	httpUrl := fmt.Sprintf("%v:%v", httpHost, httpPort)

	// convert token duration to int
	tokenDuration, err := strconv.Atoi(strTokenDuration)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// generate token
	expirationTime := time.Now().Add(time.Duration(tokenDuration) * time.Minute)

	claims := &domain.Claims{
		Username: input.Username,
		RoleID:     user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Set token string ke dalam cookie response
	maxAge := 60 * tokenDuration
	c.SetCookie("jwt_token", tokenString, maxAge, "/", httpUrl, false, true)

	// auth success: return generated token
	c.JSON(http.StatusOK, gin.H{
		"jwt_token": tokenString,
	})
}
