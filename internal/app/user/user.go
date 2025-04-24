package user

import (
	"context"
	"github.com/google/uuid"
	"time"

	"github.com/turahe/interpesona-data/internal/db/model"
	"github.com/turahe/interpesona-data/internal/repository"
	"github.com/turahe/interpesona-data/pkg/exception"
)

type UserApp interface {
	GetUsers(ctx context.Context) ([]GetUserDTO, error)
	GetUsersWithPagination(ctx context.Context, limit int, page int) ([]GetUserDTO, error)
	GetUserByID(ctx context.Context, input GetUserDTI) (GetUserDTO, error)
	CreateUser(ctx context.Context, input CreateUserDTI) (CreateUserDTO, error)
	UpdateUser(ctx context.Context, input UpdateUserDTI) (GetUserDTO, error)
	DeleteUser(ctx context.Context, input DeleteUserDTI) (bool, error)
}

type userApp struct {
	Repo *repository.Repository
}

func NewUserApp(repo *repository.Repository) UserApp {
	return &userApp{
		Repo: repo,
	}
}

type GetUserDTI struct {
	ID uuid.UUID `json:"id"`
}

type GetUserDTO struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *userApp) GetUsers(ctx context.Context) ([]GetUserDTO, error) {
	users, err := s.Repo.User.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var usersDTO []GetUserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, GetUserDTO{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email,
			Phone:    user.Phone,
		})
	}

	return usersDTO, nil
}

func (s *userApp) GetUsersWithPagination(ctx context.Context, limit int, page int) ([]GetUserDTO, error) {
	users, err := s.Repo.User.GetUsersWithPagination(ctx, limit, page)
	if err != nil {
		return nil, err
	}

	var usersDTO []GetUserDTO
	for _, user := range users {
		usersDTO = append(usersDTO, GetUserDTO{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email,
			Phone:    user.Phone,
		})
	}

	return usersDTO, nil
}

func (s *userApp) GetUserByID(ctx context.Context, input GetUserDTI) (GetUserDTO, error) {
	user, err := s.Repo.User.GetUserByID(ctx, input.ID)
	if err != nil {
		return GetUserDTO{}, err
	}
	if user.ID == uuid.Nil { // Check for zero-value UUID
		return GetUserDTO{}, exception.DataNotFoundError
	}
	return GetUserDTO{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Phone:    user.Phone,
	}, nil
}

type CreateUserDTI struct {
	UserName string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
}

type CreateUserDTO struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email" `
	Phone     string    `json:"phone" `
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *userApp) CreateUser(ctx context.Context, input CreateUserDTI) (CreateUserDTO, error) {
	// Ensure email is not already taken
	isUserEmailExist, err := s.Repo.User.IsUserEmailExist(ctx, input.Email)
	if err != nil {
		return CreateUserDTO{}, err
	}
	if isUserEmailExist {
		return CreateUserDTO{}, exception.UserEmailAlreadyTakenError
	}
	isUserPhoneExist, err := s.Repo.User.IsUserPhoneExist(ctx, input.Phone)
	if err != nil {
		return CreateUserDTO{}, err
	}
	if isUserPhoneExist {
		return CreateUserDTO{}, exception.UserPhoneAlreadyTakenError
	}

	user, err := s.Repo.User.AddUser(ctx, model.User{
		UserName: input.UserName,
		Email:    input.Email,
		Phone:    input.Phone,
	})

	if err != nil {
		return CreateUserDTO{}, err
	}

	return CreateUserDTO{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

type UpdateUserDTI struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"username" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Phone    string    `json:"phone" validate:"required"`
}

type UpdateUserDTO struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"username"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}

func (s *userApp) UpdateUser(ctx context.Context, input UpdateUserDTI) (GetUserDTO, error) {
	userRepo, err := s.Repo.User.GetUserByID(ctx, input.ID)
	if err != nil {
		return GetUserDTO{}, err
	}

	user, err := s.Repo.User.UpdateUser(ctx, model.User{
		ID:       userRepo.ID,
		UserName: input.UserName,
		Email:    input.Email,
		Phone:    input.Phone,
	})
	if err != nil {
		return GetUserDTO{}, err
	}

	if user.ID != uuid.Nil { // Check if the user was successfully updated
		return GetUserDTO{
			ID:        user.ID,
			UserName:  user.UserName,
			Email:     user.Email,
			Phone:     user.Phone,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}, nil
	}

	return GetUserDTO{}, nil
}

type DeleteUserDTI struct {
	ID uuid.UUID `json:"id"`
}

func (s *userApp) DeleteUser(ctx context.Context, input DeleteUserDTI) (bool, error) {
	user, err := s.Repo.User.GetUserByID(ctx, input.ID)
	if err != nil {
		return false, err
	}
	deleteUser, err := s.Repo.User.DeleteUser(ctx, user.ID)
	if err != nil {
		return false, err
	}

	return deleteUser, nil
}
