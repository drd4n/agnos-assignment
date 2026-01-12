package models

import "time"

type Patient struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	FirstNameTh  string    `gorm:"index:,type:btree" json:"first_name_th" validate:"required,min=1,max=100"`
	MiddleNameTh string    `json:"middle_name_th" validate:"max=100"`
	LastNameTh   string    `gorm:"index:,type:btree" json:"last_name_th" validate:"required,min=1,max=100"`
	FirstNameEn  string    `json:"first_name_en" validate:"required,min=1,max=100"`
	MiddleNameEn string    `json:"middle_name_en" validate:"max=100"`
	LastNameEn   string    `json:"last_name_en" validate:"required,min=1,max=100"`
	DateOfBirth  time.Time `json:"date_of_birth" validate:"required"`
	PatientHN    string    `gorm:"uniqueIndex" json:"patient_hn" validate:"required,min=1,max=50"`
	NationalID   string    `gorm:"uniqueIndex" json:"national_id" validate:"required,min=1,max=20"`
	PassportID   string    `gorm:"uniqueIndex" json:"passport_id" validate:"max=50"`
	PhoneNumber  string    `json:"phone_number" validate:"required,min=1,max=20"`
	Email        string    `gorm:"uniqueIndex" json:"email" validate:"required,email"`
	Gender       string    `json:"gender" validate:"required,oneof=M F"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
