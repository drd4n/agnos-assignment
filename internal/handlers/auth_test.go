package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandlerHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := &AuthHandler{}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handler.HealthCheck(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "ok", response["status"])
}

func TestAuthHandlerRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("should fail without valid request body", func(t *testing.T) {
		handler := &AuthHandler{}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/staff/create", bytes.NewReader([]byte("invalid")))
		c.Request.Header.Set("Content-Type", "application/json")

		handler.Register(c)

		assert.NotEqual(t, http.StatusCreated, w.Code)
	})
}

func TestAuthHandlerLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("should fail without valid request body", func(t *testing.T) {
		handler := &AuthHandler{}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/staff/login", bytes.NewReader([]byte("invalid")))
		c.Request.Header.Set("Content-Type", "application/json")

		handler.Login(c)

		assert.NotEqual(t, http.StatusOK, w.Code)
	})
}
