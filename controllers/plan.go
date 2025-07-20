package controllers

import (
	"net/http"

	"github.com/Ivan-S-Petkov/organisation-api/config"
	"github.com/Ivan-S-Petkov/organisation-api/models"
	"github.com/gin-gonic/gin"
)

type SwitchPlanInput struct {
  Name string `json:"name" binding:"required"`
}

var planLimits = map[string]int{
  "basic":     5,
  "pro":       25,
  "enterprise": 100,
}

func GetPlan(c *gin.Context) {
  var plan models.Plan
  if err := config.DB.First(&plan).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "Subscription plan not found"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": plan})
}

func SwitchPlan(c *gin.Context) {
  var input SwitchPlanInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Plan name required"})
    return
  }

  limit, exists := planLimits[input.Name]
  if !exists {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plan name"})
    return
  }

  var plan models.Plan
  if err := config.DB.First(&plan).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "Current plan not found"})
    return
  }

  if plan.Used > limit {
    c.JSON(http.StatusForbidden, gin.H{"error": "Too many licenses for selected plan"})
    return
  }

  plan.Name = input.Name
  plan.Limit = limit
  config.DB.Save(&plan)

  c.JSON(http.StatusOK, gin.H{"data": plan})
}