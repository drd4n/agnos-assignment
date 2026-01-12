package routes

import (
	"agnos-assignment/internal/config"
	"agnos-assignment/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config, healthHandler *handlers.HealthHandler, authHandler *handlers.AuthHandler, patientHandler *handlers.PatientHandler) {
	RegisterHealthRoutes(router, healthHandler)
	RegisterAuthRoutes(router, authHandler, cfg)
	RegisterPatientRoutes(router, patientHandler, cfg)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Route not found"})
	})
}
