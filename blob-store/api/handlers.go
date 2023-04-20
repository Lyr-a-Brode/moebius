package api

import (
	"github.com/Lyr-a-Brode/moebius/blob-store/service"
	"github.com/Lyr-a-Brode/moebius/blob-store/trace"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
)

type Handlers struct {
	storeService service.StoreService
}

func NewHandlers(storeService service.StoreService) Handlers {
	return Handlers{
		storeService: storeService,
	}
}
func (h Handlers) UploadBlob(ctx echo.Context, params UploadBlobParams) error {
	traceID := params.XTraceID.String()

	tCtx := trace.NewContext(ctx.Request().Context(), traceID)

	file, err := ctx.FormFile("file")
	if err != nil {
		log.WithError(err).WithField("trace_id", traceID).
			Error("Unable to read blob from request field")

		return ctx.JSON(http.StatusBadRequest, ErrorBadRequest{
			Code:    "blob_read_error",
			Message: "Unable to read blob from request field",
		})
	}

	format := ctx.FormValue("format")

	src, err := file.Open()
	if err != nil {
		log.WithError(err).WithField("trace_id", traceID).
			Error("Unable to open blob file")

		return ctx.JSON(http.StatusInternalServerError, ErrorInternal{
			Code:    "internal_error",
			Message: "Unable to open blob file",
		})
	}

	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			log.WithError(err).WithField("trace_id", traceID).
				Error("Unable to close blob file")
		}
	}(src)

	blobID, err := h.storeService.StoreBlob(tCtx, src, format)
	if err != nil {
		log.WithError(err).WithField("trace_id", traceID).
			Error("Unable to store blob")

		return ctx.JSON(http.StatusInternalServerError, ErrorInternal{
			Code:    "internal_error",
			Message: "Unable to store blob",
		})
	}

	return ctx.JSON(http.StatusCreated, UploadBlobSuccessResponse{BlobId: blobID})
}
