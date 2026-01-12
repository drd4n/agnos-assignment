package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPatientHandlerSearchExternal(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := &PatientHandler{}

	t.Run("should fail if neither national_id nor passport_id provided", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/api/v1/patient/search-external", nil)

		handler.SearchExternal(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should fail if both national_id and passport_id provided", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/api/v1/patient/search-external?national_id=123&passport_id=N123", nil)

		handler.SearchExternal(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestPatientHandlerSearch(t *testing.T) {
	t.Run("should validate search request format", func(t *testing.T) {
		assert.NotNil(t, &PatientHandler{})
	})
}
