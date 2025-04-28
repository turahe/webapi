package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turahe/interpesona-data/internal/app/user"
	//model "github.com/turahe/interpesona-data/internal/db/model"
	dto "github.com/turahe/interpesona-data/internal/dto"
	password "github.com/turahe/interpesona-data/internal/helper/utils"
	"github.com/turahe/interpesona-data/internal/interface/requests"
	"github.com/turahe/interpesona-data/internal/interface/response"
	"net/http"
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
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Process the business logic
	dto, err := h.app.CreateUser(c.Context(), dto.CreateUserDTI{
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: password.GeneratePassword(req.Password),
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
