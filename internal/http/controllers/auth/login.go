package auth

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"webapi/internal/app/user"
	dti "webapi/internal/dto"
	"webapi/internal/http/requests"
	"webapi/internal/http/response"
)

type LoginHTTPHandler struct {
	app user.UserApp
}

func NewLoginHTTPHandler(app user.UserApp) *LoginHTTPHandler {
	return &LoginHTTPHandler{app: app}
}

func (h *LoginHTTPHandler) Login(c *fiber.Ctx) error {
	var req requests.AuthLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Process the business logic
	dto, err := h.app.Login(c.Context(), dti.LoginUserDTI{
		UserName: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(response.CommonResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "OK",
		Data:            dto,
	})
}
