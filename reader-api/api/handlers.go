package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handlers struct{}

func NewHandlers() Handlers {
	return Handlers{}
}

func (h Handlers) UploadBook(ctx echo.Context) error {
	return ctx.NoContent(http.StatusCreated)
}
