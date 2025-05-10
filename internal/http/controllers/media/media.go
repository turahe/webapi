package media

import (
	"github.com/gofiber/fiber/v2"
	"webapi/internal/app/media"
	"webapi/internal/http/requests"
	"webapi/internal/http/response"
)

type MediaHttpHandler struct {
	app media.MediaApp
}

func NewMediaHttpHandler(app media.MediaApp) *MediaHttpHandler {
	return &MediaHttpHandler{app: app}
}

func (h *MediaHttpHandler) CreateMedia(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10) // Default to 10 if not provided
	page := c.QueryInt("page", 1)    // Default to 1 if not provided
	query := c.Query("query", "")    // Default to empty string if not provided

	offset := (page - 1) * limit
	req := requests.DataWithPaginationRequest{
		Query: query,
		Limit: limit,
		Page:  offset,
	}
	responseMedia, err := h.app.GetMediaWithPagination(c.Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(response.PaginationResponse{
		TotalCount:   responseMedia.Total,
		TotalPage:    responseMedia.Total / limit,
		CurrentPage:  page,
		LastPage:     responseMedia.LastPage,
		PerPage:      limit,
		NextPage:     page + 1,
		PreviousPage: page - 1,
		Data:         responseMedia.Data,
		Path:         c.Path(),
	})
}
