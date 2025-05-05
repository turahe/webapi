package miscellaneous

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turahe/interpesona-data/pkg/exception"
)

type MiscellaneousHTTPHandler struct{}

func NewMiscellaneousHTTPHandler() *MiscellaneousHTTPHandler {
	return &MiscellaneousHTTPHandler{}
}

func (m *MiscellaneousHTTPHandler) NotFound(c *fiber.Ctx) error {
	return exception.ApiNotFoundError
}
