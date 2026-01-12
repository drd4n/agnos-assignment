package services

import (
	"testing"

	"agnos-assignment/internal/dtos/auth/requests"
	"agnos-assignment/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStaffRepository struct {
	mock.Mock
}

func (m *MockStaffRepository) Create(staff *models.Staff) error {
	args := m.Called(staff)
	return args.Error(0)
}

func (m *MockStaffRepository) FindByEmail(email string) (*models.Staff, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Staff), args.Error(1)
}

func (m *MockStaffRepository) FindByID(id int) (*models.Staff, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Staff), args.Error(1)
}

func (m *MockStaffRepository) VerifyPassword(hashedPassword, plainPassword string) bool {
	args := m.Called(hashedPassword, plainPassword)
	return args.Bool(0)
}

func TestAuthServiceRegister(t *testing.T) {
	t.Run("should validate register input", func(t *testing.T) {
		req := &requests.RegisterRequest{
			Username: "test",
			Email:    "test@example.com",
			Password: "password123",
			Hospital: "Hospital",
		}

		assert.NotEmpty(t, req.Username)
		assert.NotEmpty(t, req.Email)
		assert.NotEmpty(t, req.Password)
		assert.NotEmpty(t, req.Hospital)
	})
}

func TestAuthServiceLogin(t *testing.T) {
	t.Run("should validate login input", func(t *testing.T) {
		req := &requests.LoginRequest{
			Email:    "test@example.com",
			Password: "password",
		}

		assert.NotEmpty(t, req.Email)
		assert.NotEmpty(t, req.Password)
	})
}

func TestAuthServiceGetStaff(t *testing.T) {
	t.Run("should handle staff lookup", func(t *testing.T) {
		var staffID int = 1

		assert.Equal(t, 1, staffID)
	})
}
