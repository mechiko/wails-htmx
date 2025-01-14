package footer

import (
	"bytes"

	"github.com/labstack/echo/v4"
)

func (t *page) Route(e *echo.Echo) error {
	e.GET("/footer", t.content)
	return nil
}

func (t *page) content(c echo.Context) error {
	var buf bytes.Buffer
	model := t.Reductor().Model()
	if err := t.Render(&buf, "content", &model, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
