package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-hexa-archi/internal/adapter/config"
	handler "github.com/yehezkiel1086/go-hexa-archi/internal/adapter/handler/http"
	"github.com/yehezkiel1086/go-hexa-archi/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-hexa-archi/internal/core/service"
)

func init() {
	config.InitEnv()
	postgres.MigrateDB()
}

func main() {
	r := gin.Default()

	// routes
  pb := r.Group("/api/v1") // public routes
	us := pb.Group("/", handler.AuthHandler()) // user routes
	ad := us.Group("/", handler.AdminHandler()) // admin routes

  // auth routes (public)
  pb.POST("/register", service.Register)
  pb.POST("/login", service.Login)

	// role routes (admin)
	ad.POST("/roles", service.CreateNewRole)

	// user routes (user)
	us.GET("/users/:user", service.GetUserByUsername)

	// user routes (admin)
	ad.GET("/users", service.GetAllUsers)

	// listen and serve on HTTP_URL:PORT (e.g: 0.0.0.0:8080)
	HTTP_HOST := os.Getenv("HTTP_HOST")
	HTTP_PORT := os.Getenv("HTTP_PORT")
	HTTP_URL := fmt.Sprintf("%v:%v", HTTP_HOST, HTTP_PORT)
  
  r.Run(HTTP_URL)
}
