package routes

import (
	"agnos-assignment/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(router *gin.Engine, healthHandler *handlers.HealthHandler) {
	router.GET("/health", healthHandler.Check)
}
