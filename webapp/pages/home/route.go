package home

import (
	"bytes"
	"firstwails/usecase"

	"github.com/labstack/echo/v4"
)

func (t *page) Route(e *echo.Echo) error {
	e.GET("/home/title", t.Title)
	e.GET("/home/ready", t.Ready)
	return nil
}

func (t *page) Title(c echo.Context) error {
	c.String(200, "Title")
	return nil
}

func (t *page) Ready(c echo.Context) error {
	if !t.reloadPage {
		// c.Response().Header().Set("HX-Refresh", "true")
		var buf bytes.Buffer
		model := t.Reductor().Model()
		model = usecase.New(t).HomeModel(model)
		if err := t.DoRender(&buf, "index", &model, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.HTML(200, buf.String())
		return nil
	}
	// без контента свап не производится
	c.NoContent(204)
	return nil
}
