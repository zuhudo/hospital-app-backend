package database

import (
	"fmt"
	"log"

	"hospital-app-backend/internal/config"
)

// DB placeholder — in production, use GORM or database/sql with PostgreSQL
type Database struct {
	Config *config.Config
}

func New(cfg *config.Config) *Database {
	return &Database{Config: cfg}
}

func (db *Database) Connect() error {
	// In production, replace with actual database connection:
	//
	// import "gorm.io/gorm"
	// import "gorm.io/driver/postgres"
	//
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	//     db.Config.DBHost, db.Config.DBUser, db.Config.DBPassword, db.Config.DBName, db.Config.DBPort)
	// gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//
	// For now, using in-memory stores in handlers

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		db.Config.DBHost, db.Config.DBUser, db.Config.DBPassword, db.Config.DBName, db.Config.DBPort)

	log.Printf("Database DSN: %s (using in-memory stores for demo)", dsn)
	return nil
}

func (db *Database) Migrate() error {
	// In production, run auto-migrations:
	// db.Conn.AutoMigrate(&models.User{}, &models.Patient{}, &models.Doctor{}, ...)
	log.Println("Database migration skipped (using in-memory stores)")
	return nil
}
