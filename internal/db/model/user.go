package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID `json:"id"`
	UserName      string    `json:"username"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Password      string    `json:"password"`
	EmailVerified time.Time `json:"email_verified"`
	PhoneVerified time.Time `json:"phone_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
