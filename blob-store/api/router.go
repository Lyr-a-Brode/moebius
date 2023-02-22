package api

import (
	"github.com/labstack/echo/v4"
	echolog "github.com/onrik/logrus/echo"
	log "github.com/sirupsen/logrus"
)

func NewRouter(handlers ServerInterface, debug bool) *echo.Echo {
	e := echo.New()

	e.Logger = echolog.NewLogger(log.StandardLogger(), "api")
	e.Debug = debug

	RegisterHandlers(e, handlers)

	return e
}
