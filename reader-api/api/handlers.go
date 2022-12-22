package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mholt/archiver"
	"github.com/nwaples/rardecode"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Handlers struct{}

func NewHandlers() Handlers {
	return Handlers{}
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

	r := &archiver.Rar{}
	err = r.Open(src, file.Size)
	if err != nil {
		return err
	}
	defer r.Close()

	for {
		f, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		th, ok := f.Header.(*rardecode.FileHeader)
		if !ok {
			return fmt.Errorf("expected header to be *rardecode.FileHeader but was %T", f.Header)
		}

		joined := filepath.Join("/tmp", th.Name)
		out, err := os.Create(joined)
		if err != nil {
			return fmt.Errorf("%s: creating new file: %v", joined, err)
		}
		defer out.Close()

		_, err = io.Copy(out, f)
		if err != nil {
			return fmt.Errorf("%s: writing file: %v", joined, err)
		}
	}

	return ctx.NoContent(http.StatusCreated)
}
