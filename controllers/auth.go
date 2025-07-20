package controllers

import (
	"net/http"

	"github.com/Ivan-S-Petkov/organisation-api/config"
	"github.com/Ivan-S-Petkov/organisation-api/models"
	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	type LoginRequest struct {
		Email string `json:"email"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}