package healthz

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"webapi/internal/http/response"
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
