package internal

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct{}

type Request struct {
	PackSizes    []int `json:"pack_sizes"`
	ItemsOrdered int   `json:"items_ordered"`
}

type Response struct {
	Result []int `json:"result"`
}

// TaskHandler
// @Summary Do Task
// @Description do task
// @ID dp-task
// @Accept  json
// @Param request body Request true "query params"
// @Produce  json
// @Success 200 {object} Response
// @Router /task [post]
func (h *Handler) TaskHandler(c echo.Context) error {
	r := new(Request)
	if err := c.Bind(r); err != nil {
		return err
	}

	result := Task(r.PackSizes, r.ItemsOrdered)

	return c.JSON(http.StatusOK, &Response{
		Result: result,
	})
}
