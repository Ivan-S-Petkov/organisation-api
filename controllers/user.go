package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Ivan-S-Petkov/organisation-api/config"
	"github.com/Ivan-S-Petkov/organisation-api/models"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
  Name  string `json:"name" binding:"required"`
  Email string `json:"email" binding:"required,email"`
  Role  string `json:"role" binding:"required,oneof=admin user"`
}

// Converts email to lowercase and trims whitespace
func normalizeEmail(email string) string {
  return strings.ToLower(strings.TrimSpace(email))
}

func CreateUser(c *gin.Context) {
  var input CreateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  user := models.User{
    Name:       input.Name, 
    Email:      normalizeEmail(input.Email), 
    Role:       input.Role,
    HasLicense: false,
  }
  
  if err := config.DB.Create(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"data": user})
}

func ListUsers(c *gin.Context) {
  var users []models.User

  query := config.DB

  if hasLicense := c.Query("hasLicense"); hasLicense != "" {
    switch hasLicense {
    case "true":
      query = query.Where("has_license = ?", true)
    case "false":
      query = query.Where("has_license = ?", false)
    }
  }

  if role := c.Query("role"); role != "" {
    query = query.Where("role = ?", role)
  }

  if search := c.Query("search"); search != "" {
    search = strings.ToLower(search)
    query = query.Where("LOWER(name) LIKE ? OR LOWER(email) LIKE ?", "%"+search+"%", "%"+search+"%")
  }

  if email := c.Query("email"); email != "" {
    query = query.Where("email = ?", normalizeEmail(email))
  }

  page := 1
  perPage := 10

  fmt.Sscanf(c.Query("page"), "%d", &page)
  fmt.Sscanf(c.Query("perPage"), "%d", &perPage)

  if page < 1 {
    page = 1
  }

  offset := (page - 1) * perPage

  var total int64
  query.Model(&models.User{}).Count(&total)

  query = query.Limit(perPage).Offset(offset)
  if err := query.Find(&users).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "data":     users,
    "page":     page,
    "perPage":  perPage,
    "total":    total,
    "pages":    (total + int64(perPage) - 1) / int64(perPage), 
  })

}
