package user

import (
	"errors"
	"github.com/google/uuid"
	dto "github.com/turahe/interpesona-data/internal/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/turahe/interpesona-data/internal/app/user"
	"github.com/turahe/interpesona-data/internal/interface/response"
	"github.com/turahe/interpesona-data/internal/interface/validation"
	"github.com/turahe/interpesona-data/pkg/exception"
)

type UserHTTPHandler struct {
	app user.UserApp
}

func NewUserHTTPHandler(app user.UserApp) *UserHTTPHandler {
	return &UserHTTPHandler{app: app}
}

// GetUsers Write me GetUsers function
func (h *UserHTTPHandler) GetUsers(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10) // Default to 10 if not provided
	page := c.QueryInt("page", 1)    // Default to 1 if not provided
	query := c.Query("query", "")    // Default to empty string if not provided

	offset := (page - 1) * limit
	req := dto.GetUsersWithPaginationDTI{
		Query: query,
		Limit: limit,
		Page:  offset,
	}
	responseUser, err := h.app.GetUsersWithPagination(c.Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(response.PaginationResponse{
		TotalCount:   responseUser.Total,
		TotalPage:    responseUser.Total / limit,
		CurrentPage:  page,
		LastPage:     responseUser.LastPage,
		PerPage:      limit,
		NextPage:     page + 1,
		PreviousPage: page - 1,
		Data:         responseUser.Data,
		Path:         c.Path(),
	})
}

func (h *UserHTTPHandler) GetUserByID(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return exception.InvalidIDError
	}

	userDto, err := h.app.GetUserByID(c.Context(), dto.GetUserDTI{ID: id})
	if err != nil {
		return err
	}

	return c.JSON(response.CommonResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "OK",
		Data:            userDto,
	})
}

func (h *UserHTTPHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserDTI

	// Parse the request body
	if err := c.BodyParser(&req); err != nil {
		return exception.InvalidRequestBodyError
	}

	// Validate the request body
	v, _ := validation.GetValidator()
	if err := v.Struct(req); err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			return exception.NewValidationFailedErrors(validationErrs)
		}
	}

	// Process the business logic
	dto, err := h.app.CreateUser(c.Context(), dto.CreateUserDTI{
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
	})

	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(response.CommonResponse{
		ResponseCode:    http.StatusCreated,
		ResponseMessage: "OK",
		Data:            dto,
	})
}

func (h *UserHTTPHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return exception.InvalidIDError
	}

	var req dto.UpdateUserDTI

	// Parse the request body
	if err := c.BodyParser(&req); err != nil {
		return exception.InvalidRequestBodyError
	}

	// Validate the request body
	v, _ := validation.GetValidator()
	if err := v.Struct(req); err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			return exception.NewValidationFailedErrors(validationErrs)
		}
	}

	dto, err := h.app.UpdateUser(c.Context(), dto.UpdateUserDTI{
		ID:       id,
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
	})

	if err != nil {
		return err
	}

	return c.JSON(response.CommonResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "OK",
		Data:            dto,
	})
}

func (h *UserHTTPHandler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return exception.InvalidIDError
	}

	_, err = h.app.DeleteUser(c.Context(), dto.DeleteUserDTI{ID: id})
	if err != nil {
		return err
	}

	return c.JSON(response.CommonResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "OK",
	})
}
