package services

import (
	"errors"
	"fmt"

	"agnos-assignment/internal/config"
	"agnos-assignment/internal/dtos"
	"agnos-assignment/internal/middleware"
	"agnos-assignment/internal/models"
	"agnos-assignment/internal/repository"
)

type AuthService struct {
	staffRepo *repository.StaffRepository
	config    *config.Config
}

func (s *AuthService) Register(req *dtos.RegisterRequest) (*dtos.AuthResponse, error) {
	existing, err := s.staffRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing staff: %w", err)
	}
	if existing != nil {
		return nil, errors.New("staff with this email already exists")
	}

	staff := &models.Staff{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Hospital: req.Hospital,
	}

	if err := s.staffRepo.Create(staff); err != nil {
		return nil, fmt.Errorf("failed to create staff: %w", err)
	}

	token, err := middleware.GenerateToken(staff.ID, staff.Email, staff.Username, staff.Hospital, s.config)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dtos.AuthResponse{
		Token:     token,
		ExpiresIn: s.config.JWT.ExpirationHours * 3600,
		Staff:     *staff,
	}, nil
}

func (s *AuthService) Login(req *dtos.LoginRequest) (*dtos.AuthResponse, error) {
	staff, err := s.staffRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find staff: %w", err)
	}
	if staff == nil {
		return nil, errors.New("invalid email or password")
	}

	if !s.staffRepo.VerifyPassword(staff.Password, req.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := middleware.GenerateToken(staff.ID, staff.Email, staff.Username, staff.Hospital, s.config)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	staff.Password = ""

	return &dtos.AuthResponse{
		Token:     token,
		ExpiresIn: s.config.JWT.ExpirationHours * 3600,
		Staff:     *staff,
	}, nil
}

func (s *AuthService) GetStaff(staffID int) (*models.Staff, error) {
	staff, err := s.staffRepo.FindByID(staffID)
	if err != nil {
		return nil, fmt.Errorf("failed to find staff: %w", err)
	}
	if staff == nil {
		return nil, errors.New("staff not found")
	}
	return staff, nil
}
