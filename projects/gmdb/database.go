package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "postgres://gmdb_user:gmdb_pass@localhost:5432/gmdb_db?sslmode=disable"
	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully!")
}

// Example models for a movie database
type Movie struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Director    string `json:"director"`
	Year        int    `json:"year"`
	Genre       string `json:"genre"`
	Rating      float64 `json:"rating"`
	Description string `json:"description"`
}

func AutoMigrate() {
	err := DB.AutoMigrate(&Movie{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration completed!")
}
