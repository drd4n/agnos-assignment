package routes

import (
	"agnos-assignment/internal/config"
	"agnos-assignment/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, authHandler *handlers.AuthHandler, cfg *config.Config) {
	staffGroup := router.Group("/staff")
	{
		staffGroup.POST("/create", authHandler.Register)
		staffGroup.POST("/login", authHandler.Login)
	}
}
