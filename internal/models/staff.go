package models

import "time"

type Staff struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"index;uniqueIndex" json:"username"`
	Email     string    `gorm:"index;uniqueIndex" json:"email"`
	Password  string    `json:"-"`
	Hospital  string    `gorm:"index" json:"hospital"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
