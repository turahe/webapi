package healthz

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turahe/interpesona-data/internal/http/response"
	"net/http"
)

type HealthzHTTPHandler struct{}

func NewHealthzHTTPHandler() *HealthzHTTPHandler {
	return &HealthzHTTPHandler{}
}

func (h *HealthzHTTPHandler) Healthz(c *fiber.Ctx) error {
	return c.JSON(response.CommonResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "OK",
	})
}
