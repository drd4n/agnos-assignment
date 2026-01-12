package responses

import "time"

type ExternalAPIResponse struct {
	FirstNameTh  string    `json:"first_name_th"`
	MiddleNameTh string    `json:"middle_name_th"`
	LastNameTh   string    `json:"last_name_th"`
	FirstNameEn  string    `json:"first_name_en"`
	MiddleNameEn string    `json:"middle_name_en"`
	LastNameEn   string    `json:"last_name_en"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	PatientHN    string    `json:"patient_hn"`
	NationalID   string    `json:"national_id"`
	PassportID   string    `json:"passport_id"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Gender       string    `json:"gender"`
}
