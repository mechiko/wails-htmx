package domain

import (
	"io"

	"github.com/labstack/echo/v4"
)

type RenderHandler func(w io.Writer, name string, data interface{}, c echo.Context) error

type EchoHandler func(c echo.Context) error
