package queue

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"webapi/internal/app/queue"
	"webapi/internal/http/response"
)

type QueueHTTPHandler struct {
	app queue.QueueApp
}

func NewQueueHTTPHandler(app queue.QueueApp) *QueueHTTPHandler {
	return &QueueHTTPHandler{app: app}
}

func (h *QueueHTTPHandler) GetQueues(c *fiber.Ctx) error {
	dtos, err := h.app.GetQueues(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(response.CommonResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "OK",
		Data:            dtos,
	})
}
