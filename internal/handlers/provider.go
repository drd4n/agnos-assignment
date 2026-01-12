package handlers

import (
	"agnos-assignment/internal/services"
)

func ProvideAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func ProvidePatientHandler(patientService *services.PatientService) *PatientHandler {
	return &PatientHandler{patientService: patientService}
}

func ProvideHealthHandler() *HealthHandler {
	return &HealthHandler{}
}
