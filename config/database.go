package config

import (
	"fmt"
	"os"

	"github.com/Ivan-S-Petkov/organisation-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
  // Build DSN from environment variables
  host := getEnv("DB_HOST", "localhost")
  user := getEnv("DB_USER", "postgres")
  password := getEnv("DB_PASSWORD", "admin")
  dbname := getEnv("DB_NAME", "organisation")
  port := getEnv("DB_PORT", "5432")
  sslmode := getEnv("DB_SSLMODE", "disable")
  timezone := getEnv("DB_TIMEZONE", "UTC")

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
    host, user, password, dbname, port, sslmode, timezone)

  database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&models.User{}, &models.Plan{})
  var existingPlan models.Plan
  if err := database.First(&existingPlan).Error; err != nil {
    basic := models.Plan{Name: "Basic", Limit: 5, Used: 0}
    database.Create(&basic)
    fmt.Println("Seeded default plan: Basic")
  }

  DB = database
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
  if value := os.Getenv(key); value != "" {
    return value
  }
  return fallback
}