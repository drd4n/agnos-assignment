package services

import (
	"net/http"
	"time"

	"agnos-assignment/internal/config"
	"agnos-assignment/internal/repository"
)

func ProvideAuthService(staffRepo *repository.StaffRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		staffRepo: staffRepo,
		config:    cfg,
	}
}

func ProvidePatientService(patientRepo *repository.PatientRepository, cfg *config.Config) *PatientService {
	return &PatientService{
		repo:   patientRepo,
		apiURL: cfg.Server.ExternalAPIURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
