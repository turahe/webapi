package user

import (
	"errors"
	"github.com/google/uuid"
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

// Write me GetUsers function
func (h *UserHTTPHandler) GetUsers(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10) // Default to 10 if not provided
	page := c.QueryInt("page", 1)    // Default to 1 if not provided
	dtos, err := h.app.GetUsersWithPagination(c.Context(), limit, page)
	if err != nil {
		return err
	}

	return c.JSON(response.CommonResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "OK",
		Data:            dtos,
		Path:            c.Path(),
	})
}

func (h *UserHTTPHandler) GetUserByID(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return exception.InvalidIDError
	}

	dti := user.GetUserDTI{ID: id}
	dto, err := h.app.GetUserByID(c.Context(), dti)
	if err != nil {
		return err
	}

	return c.JSON(response.CommonResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "OK",
		Data:            dto,
	})
}

func (h *UserHTTPHandler) CreateUser(c *fiber.Ctx) error {
	var req user.CreateUserDTI

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
	dto, err := h.app.CreateUser(c.Context(), user.CreateUserDTI{
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

	var req user.UpdateUserDTI

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

	dto, err := h.app.UpdateUser(c.Context(), user.UpdateUserDTI{
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

	_, err = h.app.DeleteUser(c.Context(), user.DeleteUserDTI{ID: id})
	if err != nil {
		return err
	}

	return c.JSON(response.CommonResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "OK",
	})
}
