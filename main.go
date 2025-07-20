package main

import (
	"log"
	"os"

	"github.com/Ivan-S-Petkov/organisation-api/config"
	"github.com/Ivan-S-Petkov/organisation-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config.ConnectDatabase()

	r := gin.Default()

	// Configure CORS middleware
	corsOrigin := getEnv("CORS_ORIGIN", "http://localhost:3000")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{corsOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-User-ID", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
  
	routes.RegisterRoutes(r)
	
	port := getEnv("PORT", "8080")
	r.Run(":" + port)
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
