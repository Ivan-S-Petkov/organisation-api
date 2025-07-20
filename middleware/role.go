package middleware

import (
	"net/http"

	"github.com/Ivan-S-Petkov/organisation-api/config"
	"github.com/Ivan-S-Petkov/organisation-api/models"
	"github.com/gin-gonic/gin"
)

func RequireAdmin() gin.HandlerFunc {
  return func(c *gin.Context) {
    userID := c.GetHeader("X-User-ID")

    var user models.User
    if err := config.DB.First(&user, userID).Error; err != nil || user.Role != "admin" {
      c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
      return
    }

    c.Next()
  }
}
