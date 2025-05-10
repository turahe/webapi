package user

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"webapi/internal/http/requests"
	"webapi/internal/http/response"
	"webapi/internal/http/validation"
	"webapi/pkg/exception"
)

func (h *UserHTTPHandler) ChangePassword(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return exception.InvalidIDError
	}
	var req requests.ChangePasswordRequest
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
	user, err := h.app.GetUserByID(c.Context(), requests.GetUserIdRequest{ID: userID})
	if err != nil {
		return exception.DataNotFoundError
	}
	dto, err := h.app.ChangePassword(c.Context(), user.ID, requests.ChangePasswordRequest{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
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

func (h *UserHTTPHandler) ChangeUserName(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userId, err := uuid.Parse(idParam)
	if err != nil {
		return exception.InvalidIDError
	}
	var req requests.ChangeUserNameRequest
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
	user, err := h.app.GetUserByID(c.Context(), requests.GetUserIdRequest{ID: userId})
	if err != nil {
		return exception.DataNotFoundError
	}
	dto, err := h.app.ChangeUserName(c.Context(), user.ID, requests.ChangeUserNameRequest{
		UserName: req.UserName,
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

func (h *UserHTTPHandler) ChangePhone(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return exception.InvalidIDError
	}
	var req requests.ChangePhoneRequest
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
	user, err := h.app.GetUserByID(c.Context(), requests.GetUserIdRequest{ID: userID})
	if err != nil {
		return exception.DataNotFoundError
	}
	dto, err := h.app.ChangePhone(c.Context(), user.ID, requests.ChangePhoneRequest{
		Phone: req.Phone,
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
func (h *UserHTTPHandler) ChangeEmail(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return exception.InvalidIDError
	}
	var req requests.ChangeEmailRequest
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
	user, err := h.app.GetUserByID(c.Context(), requests.GetUserIdRequest{ID: userID})
	if err != nil {
		return exception.DataNotFoundError
	}
	dto, err := h.app.ChangeEmail(c.Context(), user.ID, requests.ChangeEmailRequest{
		Email: req.Email,
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
