package header

import (
	"bytes"
	"firstwails/usecase"

	"github.com/labstack/echo/v4"
)

func (t *page) Route(e *echo.Echo) error {
	e.GET("/header", t.content)
	e.GET("/header/modal", t.mainmenu)
	return nil
}

func (t *page) mainmenu(c echo.Context) error {
	var buf bytes.Buffer
	model := t.Reductor().Model()
	model = usecase.New(t).DbInfoModel(model)
	if err := t.Render(&buf, "modal", &model); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}

func (t *page) content(c echo.Context) error {
	var buf bytes.Buffer
	model := t.Reductor().Model()
	model = usecase.New(t).DbInfoModel(model)
	if err := t.Render(&buf, "content", &model); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
