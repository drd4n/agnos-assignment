package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"agnos-assignment/internal/dtos"
	"agnos-assignment/internal/models"
	"agnos-assignment/internal/repository"
)

type PatientService struct {
	repo       *repository.PatientRepository
	apiURL     string
	httpClient *http.Client
}

func (s *PatientService) GetPatientExternal(identifier string) (*models.Patient, error) {
	url := fmt.Sprintf("%s/patient/search/%s", s.apiURL, identifier)

	resp, err := s.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call external API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp dtos.ExternalAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	var existingPatient *models.Patient

	if apiResp.NationalID != "" {
		existingPatient, _ = s.repo.FindByNationalID(apiResp.NationalID)
	}

	if existingPatient == nil && apiResp.PassportID != "" {
		existingPatient, _ = s.repo.FindByPassportID(apiResp.PassportID)
	}

	patient := &models.Patient{
		FirstNameTh:  apiResp.FirstNameTh,
		MiddleNameTh: apiResp.MiddleNameTh,
		LastNameTh:   apiResp.LastNameTh,
		FirstNameEn:  apiResp.FirstNameEn,
		MiddleNameEn: apiResp.MiddleNameEn,
		LastNameEn:   apiResp.LastNameEn,
		DateOfBirth:  apiResp.DateOfBirth,
		PatientHN:    apiResp.PatientHN,
		NationalID:   apiResp.NationalID,
		PassportID:   apiResp.PassportID,
		PhoneNumber:  apiResp.PhoneNumber,
		Email:        apiResp.Email,
		Gender:       apiResp.Gender,
	}

	if existingPatient != nil {
		patient.ID = existingPatient.ID
		if err := s.repo.Update(patient); err != nil {
			return nil, fmt.Errorf("failed to update patient: %w", err)
		}
	} else {
		if err := s.repo.Create(patient); err != nil {
			return nil, fmt.Errorf("failed to create patient: %w", err)
		}
	}

	return patient, nil
}

func (s *PatientService) SearchPatients(ctx context.Context, criteria *dtos.PatientSearchCriteria) ([]*models.Patient, error) {
	return s.repo.SearchPatients(ctx, criteria)
}
