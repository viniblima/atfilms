package database

import (
	"fmt"
	"log"
	"os"

	"github.com/viniblima/atfilms/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func ConnectDb() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(
		&models.User{},
		&models.Customer{},
		&models.Photo{},
		&models.Video{},
		&models.Job{},
	)

	newUser := models.User{
		Name:     os.Getenv("SUPERUSER_NAME"),
		Email:    os.Getenv("SUPERUSER_EMAIL"),
		Password: os.Getenv("SUPERUSER_PASSWORD"),
	}

	db.Create(&newUser)

	Db = db
}
