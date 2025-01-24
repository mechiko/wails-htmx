package setup

import (
	"github.com/labstack/echo/v4"
)

func (t *page) Route(e *echo.Echo) error {
	e.GET("/setup/title", t.Title)
	e.POST("/setup/omsid", t.ValidateOmsID)
	e.POST("/setup/deviceid", t.ValidateDeviceID)
	e.POST("/setup/configdb", t.ValidateConfigDB)
	e.POST("/setup/save", t.Save)
	e.GET("/setup/ready", t.Ready)
	e.POST("/setup/ping", t.Ping)
	return nil
}

func (t *page) Title(c echo.Context) error {
	c.String(200, "Title")
	return nil
}
