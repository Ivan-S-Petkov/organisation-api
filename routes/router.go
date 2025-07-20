package routes

import (
	"github.com/Ivan-S-Petkov/organisation-api/controllers"
	"github.com/Ivan-S-Petkov/organisation-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
  r.POST("/login", controllers.LoginUser)

  r.POST("/users", middleware.RequireAdmin(), controllers.CreateUser)
  r.GET("/users", controllers.ListUsers)

  r.GET("/plan", controllers.GetPlan)
  r.POST("/plan/switch", middleware.RequireAdmin(), controllers.SwitchPlan)


  r.POST("/licenses/assign", middleware.RequireAdmin(), controllers.AssignLicense)
  r.POST("/licenses/unassign", middleware.RequireAdmin(), controllers.UnassignLicense)
}