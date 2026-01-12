package routes

import (
	"agnos-assignment/internal/config"
	"agnos-assignment/internal/handlers"
	"agnos-assignment/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPatientRoutes(router *gin.Engine, patientHandler *handlers.PatientHandler, cfg *config.Config) {
	patientGroup := router.Group("/api/v1/patient")
	patientGroup.Use(middleware.AuthMiddleware(cfg))
	{
		patientGroup.POST("/search", patientHandler.Search)
		patientGroup.GET("/search-external", patientHandler.SearchExternal)
	}
}
