package user

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
type GetUsersWithPaginationDTO struct {
	Total       int          `json:"total"`
	Limit       int          `json:"limit"`
	CurrentPage int          `json:"currentPage"`
	LastPage    int          `json:"lastPage"`
	Data        []GetUserDTO `json:"data"`
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

type GetUsersWithPaginationDTI struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}
