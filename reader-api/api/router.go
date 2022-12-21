package api

import (
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echolog "github.com/onrik/logrus/echo"
	log "github.com/sirupsen/logrus"
)

func NewRouter(handlers ServerInterface) (e *echo.Echo, err error) {
	swagger, err := GetSwagger()
	if err != nil {
		return nil, err
	}

	swagger.Servers = nil

	e = echo.New()

	e.Logger = echolog.NewLogger(log.StandardLogger(), "api")
	e.Use(middleware.OapiRequestValidator(swagger))

	RegisterHandlers(e, handlers)

	return e, nil
}
