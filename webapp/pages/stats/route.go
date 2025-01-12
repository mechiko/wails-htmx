package stats

import (
	"bytes"
	"firstwails/webapp/htmxutil"

	"github.com/labstack/echo/v4"
)

func (t *page) Route(e *echo.Echo) error {
	e.GET("/stats/title", t.Title)
	e.GET("/stats/ready", t.Ready)
	e.GET("/stats/modal", t.modal)
	return nil
}

func (t *page) Title(c echo.Context) error {
	c.String(200, "STATS")
	return nil
}

func (t *page) Ready(c echo.Context) error {
	if t.ReloadActivePage() {
		t.SetReloadActivePage(false)
		var buf bytes.Buffer
		model := t.Reductor().Model()
		if err := t.Render(&buf, "index", &model, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.HTML(200, buf.String())
		return nil
	}
	hx := htmxutil.HxRequestHeaderFromRequest(c.Request())
	if !hx.HxRequest {
		// готовность DOM когда не надо обновлять страницу и прилетел пинг от нее
		t.DomReadyHttp()
	}
	// без контента свап не производится
	c.NoContent(204)
	return nil
}

func (t *page) modal(c echo.Context) error {
	var buf bytes.Buffer
	model := t.Reductor().Model()
	if err := t.Render(&buf, "modal", &model, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
