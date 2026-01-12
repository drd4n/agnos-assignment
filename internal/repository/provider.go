package repository

import (
	"agnos-assignment/internal/database"
)

func ProvideStaffRepository(db *database.Connection) *StaffRepository {
	return &StaffRepository{db: db.DB}
}

func ProvidePatientRepository(db *database.Connection) *PatientRepository {
	return &PatientRepository{db: db.DB}
}
