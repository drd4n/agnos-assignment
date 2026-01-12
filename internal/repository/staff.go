package repository

import (
	"agnos-assignment/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type StaffRepository struct {
	db *gorm.DB
}

func (r *StaffRepository) Create(staff *models.Staff) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(staff.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	staff.Password = string(hashedPassword)

	if result := r.db.Create(staff); result.Error != nil {
		return result.Error
	}

	staff.Password = ""
	return nil
}

func (r *StaffRepository) FindByEmail(email string) (*models.Staff, error) {
	staff := &models.Staff{}

	if result := r.db.Where("email = ?", email).First(staff); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return staff, nil
}

func (r *StaffRepository) FindByID(id int) (*models.Staff, error) {
	staff := &models.Staff{}

	if result := r.db.First(staff, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return staff, nil
}

func (r *StaffRepository) VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
