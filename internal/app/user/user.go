package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	user "webapi/internal/dto"
	password "webapi/internal/helper/utils"

	"webapi/internal/db/model"
	"webapi/internal/repository"
	"webapi/pkg/exception"
)

type UserApp interface {
	Login(ctx context.Context, dti user.LoginUserDTI) (user.GetUserDTO, error)
	GetUsers(ctx context.Context) ([]user.GetUserDTO, error)
	GetUsersWithPagination(ctx context.Context, input user.GetUsersWithPaginationDTI) (user.GetUsersWithPaginationDTO, error)
	GetUserByID(ctx context.Context, input user.GetUserDTI) (user.GetUserDTO, error)
	CreateUser(ctx context.Context, input user.CreateUserDTI) (user.GetUserDTO, error)
	UpdateUser(ctx context.Context, input user.UpdateUserDTI) (user.GetUserDTO, error)
	DeleteUser(ctx context.Context, input user.DeleteUserDTI) (bool, error)
}

type userApp struct {
	Repo *repository.Repository
}

func NewUserApp(repo *repository.Repository) UserApp {
	return &userApp{
		Repo: repo,
	}
}

func (s *userApp) Login(ctx context.Context, dti user.LoginUserDTI) (user.GetUserDTO, error) {
	userRepo, err := s.Repo.User.GetUserByUsername(ctx, dti.UserName)
	if err != nil {
		return user.GetUserDTO{}, err
	}
	fmt.Println(userRepo)
	authLogin := password.ComparePassword(userRepo.Password, dti.Password)

	if userRepo.ID == uuid.Nil { // Check for zero-value UUID
		return user.GetUserDTO{}, exception.DataNotFoundError
	}
	if authLogin {
		return user.GetUserDTO{
			ID:        userRepo.ID,
			UserName:  userRepo.UserName,
			Email:     userRepo.Email,
			Phone:     userRepo.Phone,
			CreatedAt: userRepo.CreatedAt,
			UpdatedAt: userRepo.UpdatedAt,
		}, nil

	}
	return user.GetUserDTO{}, exception.InvalidCredentialsError

}

func (s *userApp) GetUsers(ctx context.Context) ([]user.GetUserDTO, error) {
	users, err := s.Repo.User.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var usersDTO []user.GetUserDTO
	for _, userRepo := range users {
		usersDTO = append(usersDTO, user.GetUserDTO{
			ID:       userRepo.ID,
			UserName: userRepo.UserName,
			Email:    userRepo.Email,
			Phone:    userRepo.Phone,
		})
	}

	return usersDTO, nil
}

func (s *userApp) GetUsersWithPagination(ctx context.Context, input user.GetUsersWithPaginationDTI) (user.GetUsersWithPaginationDTO, error) {
	responseUser, err := s.Repo.User.GetUsersWithPagination(ctx, input)
	if err != nil {
		return user.GetUsersWithPaginationDTO{}, err
	}

	return responseUser, nil
}

func (s *userApp) GetUserByID(ctx context.Context, input user.GetUserDTI) (user.GetUserDTO, error) {
	userRepo, err := s.Repo.User.GetUserByID(ctx, input)
	if err != nil {
		return user.GetUserDTO{}, err
	}
	if userRepo.ID == uuid.Nil { // Check for zero-value UUID
		return user.GetUserDTO{}, exception.DataNotFoundError
	}
	return user.GetUserDTO{
		ID:       userRepo.ID,
		UserName: userRepo.UserName,
		Email:    userRepo.Email,
		Phone:    userRepo.Phone,
	}, nil
}

func (s *userApp) CreateUser(ctx context.Context, input user.CreateUserDTI) (user.GetUserDTO, error) {
	// Ensure email is not already taken
	isUserEmailExist, err := s.Repo.User.IsUserEmailExist(ctx, input.Email)
	if err != nil {
		return user.GetUserDTO{}, err
	}
	if isUserEmailExist {
		return user.GetUserDTO{}, exception.UserEmailAlreadyTakenError
	}
	isUserPhoneExist, err := s.Repo.User.IsUserPhoneExist(ctx, input.Phone)
	if err != nil {
		return user.GetUserDTO{}, err
	}
	if isUserPhoneExist {
		return user.GetUserDTO{}, exception.UserPhoneAlreadyTakenError
	}

	userRepo, err := s.Repo.User.AddUser(ctx, model.User{
		UserName: input.UserName,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: password.GeneratePassword(input.Password),
	})

	if err != nil {
		return user.GetUserDTO{}, err
	}

	return user.GetUserDTO{
		ID:        userRepo.ID,
		UserName:  userRepo.UserName,
		Email:     userRepo.Email,
		Phone:     userRepo.Phone,
		CreatedAt: userRepo.CreatedAt,
		UpdatedAt: userRepo.UpdatedAt,
	}, nil
}

func (s *userApp) UpdateUser(ctx context.Context, input user.UpdateUserDTI) (user.GetUserDTO, error) {
	id := user.GetUserDTI{
		ID: input.ID,
	}
	userRepo, err := s.Repo.User.GetUserByID(ctx, id)
	if err != nil {
		return user.GetUserDTO{}, err
	}

	userRepo, err = s.Repo.User.UpdateUser(ctx, model.User{
		ID:       userRepo.ID,
		UserName: input.UserName,
		Email:    input.Email,
		Phone:    input.Phone,
	})
	if err != nil {
		return user.GetUserDTO{}, err
	}

	if userRepo.ID != uuid.Nil { // Check if the user was successfully updated
		return user.GetUserDTO{
			ID:        userRepo.ID,
			UserName:  userRepo.UserName,
			Email:     userRepo.Email,
			Phone:     userRepo.Phone,
			CreatedAt: userRepo.CreatedAt,
			UpdatedAt: userRepo.UpdatedAt,
		}, nil
	}

	return user.GetUserDTO{}, nil
}

func (s *userApp) DeleteUser(ctx context.Context, input user.DeleteUserDTI) (bool, error) {
	userRepo, err := s.Repo.User.GetUserByID(ctx, user.GetUserDTI{ID: input.ID})
	if err != nil {
		return false, err
	}
	deleteUser, err := s.Repo.User.DeleteUser(ctx, user.GetUserDTI{ID: userRepo.ID})
	if err != nil {
		return false, err
	}

	return deleteUser, nil
}
