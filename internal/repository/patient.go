package repository

import (
	"context"
	"errors"
	"fmt"

	"agnos-assignment/internal/dtos"
	"agnos-assignment/internal/models"

	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

func (r *PatientRepository) Create(patient *models.Patient) error {
	if result := r.db.Create(patient); result.Error != nil {
		return fmt.Errorf("failed to create patient: %w", result.Error)
	}

	return nil
}

func (r *PatientRepository) FindByID(id int) (*models.Patient, error) {
	patient := &models.Patient{}
	if result := r.db.First(patient, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("patient not found")
		}
		return nil, fmt.Errorf("failed to find patient: %w", result.Error)
	}

	return patient, nil
}

func (r *PatientRepository) FindByPatientHN(patientHN string) (*models.Patient, error) {
	patient := &models.Patient{}
	if result := r.db.Where("patient_hn = ?", patientHN).First(patient); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("patient not found")
		}
		return nil, fmt.Errorf("failed to find patient: %w", result.Error)
	}

	return patient, nil
}

func (r *PatientRepository) Update(patient *models.Patient) error {
	result := r.db.Save(patient)
	if result.Error != nil {
		return fmt.Errorf("failed to update patient: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("patient not found")
	}

	return nil
}

func (r *PatientRepository) Delete(id int) error {
	result := r.db.Delete(&models.Patient{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete patient: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("patient not found")
	}

	return nil
}

func (r *PatientRepository) GetAll() ([]*models.Patient, error) {
	var patients []*models.Patient

	if result := r.db.Order("created_at DESC").Find(&patients); result.Error != nil {
		return nil, fmt.Errorf("failed to get patients: %w", result.Error)
	}

	return patients, nil
}

func (r *PatientRepository) FindByNationalID(nationalID string) (*models.Patient, error) {
	patient := &models.Patient{}
	if result := r.db.Where("national_id = ?", nationalID).First(patient); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("patient not found")
		}
		return nil, fmt.Errorf("failed to find patient: %w", result.Error)
	}

	return patient, nil
}

func (r *PatientRepository) FindByPassportID(passportID string) (*models.Patient, error) {
	patient := &models.Patient{}
	if result := r.db.Where("passport_id = ?", passportID).First(patient); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("patient not found")
		}
		return nil, fmt.Errorf("failed to find patient: %w", result.Error)
	}

	return patient, nil
}

func (r *PatientRepository) SearchPatients(ctx context.Context, criteria *dtos.PatientSearchCriteria) ([]*models.Patient, error) {
	var patients []*models.Patient
	query := r.db
	query = query.Where("patient_hn = ?", ctx.Value("hospital"))

	if criteria.NationalID != "" {
		query = query.Where("national_id = ?", criteria.NationalID)
	}

	if criteria.PassportID != "" {
		query = query.Where("passport_id = ?", criteria.PassportID)
	}

	if criteria.FirstName != "" {
		query = query.Where("first_name_th ILIKE ? OR first_name_en ILIKE ?", "%"+criteria.FirstName+"%", "%"+criteria.FirstName+"%")
	}

	if criteria.MiddleName != "" {
		query = query.Where("middle_name_th ILIKE ? OR middle_name_en ILIKE ?", "%"+criteria.MiddleName+"%", "%"+criteria.MiddleName+"%")
	}

	if criteria.LastName != "" {
		query = query.Where("last_name_th ILIKE ? OR last_name_en ILIKE ?", "%"+criteria.LastName+"%", "%"+criteria.LastName+"%")
	}

	if !criteria.DateOfBirth.IsZero() {
		query = query.Where("date_of_birth = ?", criteria.DateOfBirth)
	}

	if criteria.PhoneNumber != "" {
		query = query.Where("phone_number ILIKE ?", "%"+criteria.PhoneNumber+"%")
	}

	if criteria.Email != "" {
		query = query.Where("email ILIKE ?", "%"+criteria.Email+"%")
	}

	if result := query.Order("created_at DESC").Find(&patients); result.Error != nil {
		return nil, fmt.Errorf("failed to search patients: %w", result.Error)
	}

	return patients, nil
}
