package auth

import (
	"github.com/gofiber/fiber/v2"
	"webapi/internal/app/user"
	//model "webapi/internal/db/model"
	"net/http"
	"webapi/internal/http/requests"
	"webapi/internal/http/response"
)

type RegisterHTTPHandler struct {
	app user.UserApp
}

func NewRegisterHTTPHandler(app user.UserApp) *RegisterHTTPHandler {
	return &RegisterHTTPHandler{app: app}
}

func (h *RegisterHTTPHandler) Register(c *fiber.Ctx) error {
	var req requests.AuthRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Process the business logic
	dto, err := h.app.CreateUser(c.Context(), requests.CreateUserRequest{
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
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
