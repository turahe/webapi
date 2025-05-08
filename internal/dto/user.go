package dto

import (
	"github.com/google/uuid"
	"time"
)

type GetUserDTO struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateUserDTO struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"username"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}

type UpdateUserDTI struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"username"`
	Email    string    `json:"email" validate:"required,email"`
	Phone    string    `json:"phone"`
}

type GetUserDTI struct {
	ID uuid.UUID `json:"id"`
}

type CreateUserDTI struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type CreateUserDTO struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email" `
	Phone     string    `json:"phone" `
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DeleteUserDTI struct {
	ID uuid.UUID `json:"id"`
}

type LoginUserDTI struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
