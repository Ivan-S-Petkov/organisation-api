package controllers

import (
	"net/http"

	"github.com/Ivan-S-Petkov/organisation-api/config"
	"github.com/Ivan-S-Petkov/organisation-api/models"
	"github.com/gin-gonic/gin"
)

type LicenseInput struct {
  UserID uint `json:"userId" binding:"required"`
}

func AssignLicense(c *gin.Context) {
  var input LicenseInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User ID required"})
    return
  }

  var user models.User
  if err := config.DB.First(&user, input.UserID).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    return
  }

  if user.HasLicense {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User already has a license"})
    return
  }

  var plan models.Plan
  if err := config.DB.First(&plan).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Plan lookup failed"})
    return
  }

  if plan.Used >= plan.Limit {
    c.JSON(http.StatusForbidden, gin.H{"error": "License limit reached"})
    return
  }

  user.HasLicense = true
  plan.Used += 1
  config.DB.Save(&user)
  config.DB.Save(&plan)

  c.JSON(http.StatusOK, gin.H{"message": "License assigned"})
}

func UnassignLicense(c *gin.Context) {
  var input LicenseInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User ID required"})
    return
  }

  var user models.User
  if err := config.DB.First(&user, input.UserID).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    return
  }

  if !user.HasLicense {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User has no license to unassign"})
    return
  }

  var plan models.Plan
  if err := config.DB.First(&plan).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Plan lookup failed"})
    return
  }

  user.HasLicense = false
  plan.Used -= 1
  config.DB.Save(&user)
  config.DB.Save(&plan)

  c.JSON(http.StatusOK, gin.H{"message": "License unassigned"})
}