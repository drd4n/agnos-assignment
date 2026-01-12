package handlers

import (
	"net/http"
	"time"

	"agnos-assignment/internal/dtos"
	"agnos-assignment/internal/middleware"
	"agnos-assignment/internal/services"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	patientService *services.PatientService
}

type SearchPatientRequest struct {
	NationalID  string    `json:"national_id"`
	PassportID  string    `json:"passport_id"`
	FirstName   string    `json:"first_name"`
	MiddleName  string    `json:"middle_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
}

func (h *PatientHandler) SearchExternal(c *gin.Context) {
	nationalID, passportID := c.Query("national_id"), c.Query("passport_id")

	if (nationalID == "" && passportID == "") || (nationalID != "" && passportID != "") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "national_id or passport_id is required"})
		return
	}

	id := nationalID
	if id == "" {
		id = passportID
	}

	response, err := h.patientService.GetPatientExternal(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *PatientHandler) Search(c *gin.Context) {
	var req SearchPatientRequest

	if c.Request.ContentLength > 0 {
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
	}

	req.NationalID = c.DefaultQuery("national_id", req.NationalID)
	req.PassportID = c.DefaultQuery("passport_id", req.PassportID)
	req.FirstName = c.DefaultQuery("first_name", req.FirstName)
	req.MiddleName = c.DefaultQuery("middle_name", req.MiddleName)
	req.LastName = c.DefaultQuery("last_name", req.LastName)
	req.PhoneNumber = c.DefaultQuery("phone_number", req.PhoneNumber)
	req.Email = c.DefaultQuery("email", req.Email)

	if dobStr := c.DefaultQuery("date_of_birth", ""); dobStr != "" {
		if dob, err := time.Parse("2006-01-02", dobStr); err == nil {
			req.DateOfBirth = dob
		}
	}

	ctx := middleware.GetContextFromGin(c)
	patients, err := h.patientService.SearchPatients(ctx, &dtos.PatientSearchCriteria{
		NationalID:  req.NationalID,
		PassportID:  req.PassportID,
		FirstName:   req.FirstName,
		MiddleName:  req.MiddleName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(patients) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})
}
