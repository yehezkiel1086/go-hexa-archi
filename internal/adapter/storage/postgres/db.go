package postgres

import (
	"fmt"
	"log"

	"github.com/yehezkiel1086/go-hexa-archi/internal/adapter/config"
	"github.com/yehezkiel1086/go-hexa-archi/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateDB() {
	// get .env DB values
	dsn := config.GetDBEnv()
	
	// connect DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Enable uuid-ossp extension (safe even if already exists)
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		log.Fatal("Failed to enable uuid-ossp extension:", err)
	}

	fmt.Println("Connected to DB!")

	// Migrate DB schemas
	db.AutoMigrate(&domain.User{}, &domain.Role{})
}

func ConnDB() (*gorm.DB, error) {
	// get .env DB values
	dsn := config.GetDBEnv()
	
	// connect DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
