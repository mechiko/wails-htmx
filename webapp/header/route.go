package header

import (
	"bytes"
	"firstwails/domain"
	"firstwails/usecase"

	"github.com/labstack/echo/v4"
)

func (t *page) Route(e *echo.Echo) error {
	e.GET("/header", t.content)
	e.GET("/header/modal", t.modal)
	e.GET("/header/:page", t.pager)
	return nil
}

func (t *page) modal(c echo.Context) error {
	var buf bytes.Buffer
	model := t.Reductor().Model()
	model = usecase.New(t).MenuModel(model)
	if err := t.Render(&buf, "modal", &model.Menu); err != nil {
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
	if err := t.Render(&buf, "content", &model.Menu); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}

func (t *page) pager(c echo.Context) error {
	page := c.Param("page")
	t.SetActivePage(page)
	// t.SetReloadActivePage(true)
	model := t.Reductor().Model()
	msg := domain.Message{
		Sender: "header.pager",
		Cmd:    page,
		Model:  &model,
	}
	t.Effects().ChanIn() <- msg

	c.Response().Header().Set("HX-Refresh", "true")
	c.NoContent(204)
	return nil
}
