package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Database *gorm.DB

func Connect() {
	log.Info("Setting up new database connection")
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := 5432
	dbSSLMode := "disable"

	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s  sslmode=%s", dbHost, dbPort, dbUser, dbName, dbPassword, dbSSLMode)
	Database, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Error("Error connecting to database")
	}
	log.Info("Database connection established")
}
