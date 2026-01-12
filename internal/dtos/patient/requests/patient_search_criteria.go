package requests

import "time"

type PatientSearchCriteria struct {
	NationalID  string
	PassportID  string
	FirstName   string
	MiddleName  string
	LastName    string
	DateOfBirth time.Time
	PhoneNumber string
	Email       string
}
