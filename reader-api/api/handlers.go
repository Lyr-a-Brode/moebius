package api

import (
	"github.com/Lyr-a-Brode/moebius/reader-api/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handlers struct {
	bookService services.BookService
}

func NewHandlers(bookService services.BookService) Handlers {
	return Handlers{
		bookService: bookService,
	}
}

func (h Handlers) UploadBook(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	err = h.bookService.ProcessBook(ctx.Request().Context(), src, file.Size)

	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}
