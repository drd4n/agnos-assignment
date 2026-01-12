package dtos

import "agnos-assignment/internal/models"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Hospital string `json:"hospital" validate:"required,min=2,max=100"`
}

type AuthResponse struct {
	Token     string       `json:"token"`
	ExpiresIn int          `json:"expires_in"`
	Staff     models.Staff `json:"staff"`
}
