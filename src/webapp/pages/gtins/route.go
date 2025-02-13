package gtins

import (
	"bytes"
	"firstwails/webapp/htmxutil"

	"github.com/labstack/echo/v4"
)

func (t *page) Route(e *echo.Echo) error {
	e.GET("/gtins/title", t.Title)
	e.GET("/gtins/ready", t.Ready)
	e.GET("/gtins/modal", t.modal)
	e.POST("/gtins/upload", t.upload)
	e.GET("/stats/progress", t.progress)
	return nil
}

func (t *page) Title(c echo.Context) error {
	c.String(200, "GTINS")
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

func (t *page) reset(c echo.Context) error {
	var buf bytes.Buffer
	// обновляем модель Stats
	model := t.Reductor().Model()
	t.UpdateModel(model, "gtins.reset")
	if err := t.Render(&buf, "page", &model, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}

func (t *page) progress(c echo.Context) error {
	var buf bytes.Buffer
	// обновляем модель Stats
	model := t.Reductor().Model().Stats
	if model.State == 1 {
		if err := t.Render(&buf, "progress", &model, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.String(200, buf.String())
	}
	c.String(200, "")
	return nil
}
