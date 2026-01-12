package responses

import "agnos-assignment/internal/models"

type AuthResponse struct {
	Token     string       `json:"token"`
	ExpiresIn int          `json:"expires_in"`
	Staff     models.Staff `json:"staff"`
}
