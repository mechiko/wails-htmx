package home

import "github.com/labstack/echo/v4"

func (t *page) Route(e *echo.Echo) error {
	e.GET("/home/title", t.Title)
	return nil
}

func (t *page) Title(c echo.Context) error {
	c.String(200, "Title")
	return nil
}
