package handlers

import (
	"net/http"

	"agnos-assignment/internal/dtos/auth/requests"
	"agnos-assignment/internal/middleware"
	"agnos-assignment/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req requests.RegisterRequest

	if !middleware.ValidateRequest(c, &req) {
		return
	}

	resp, err := h.authService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req requests.LoginRequest

	if !middleware.ValidateRequest(c, &req) {
		return
	}

	resp, err := h.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
