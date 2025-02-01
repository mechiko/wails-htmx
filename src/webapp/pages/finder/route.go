package finder

import (
	"bytes"
	"firstwails/webapp/htmxutil"
	"fmt"

	"github.com/labstack/echo/v4"
)

func (t *page) Route(e *echo.Echo) error {
	e.GET("/finder/title", t.Title)
	e.GET("/finder/ready", t.Ready)
	e.GET("/finder/modal", t.modal)
	e.POST("/finder/upload", t.upload)
	e.GET("/finder/progress", t.progress)
	e.GET("/finder/datamatrix", t.datamatrix)
	e.GET("/finder/datamatrixlist", t.datamatrixList)
	return nil
}

func (t *page) Title(c echo.Context) error {
	c.String(200, "Finder")
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
	t.UpdateModel(model, "finder.reset")
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

func (t *page) datamatrix(c echo.Context) error {
	port := t.Configuration().HostPort
	host := t.Configuration().Hostname
	proto := `http://`
	url := fmt.Sprintf("%s%s:%s/finder/datamatrixlist", proto, host, port)
	t.Open(url)
	c.NoContent(204)
	return nil
}

func (t *page) datamatrixList(c echo.Context) error {
	var buf bytes.Buffer
	model := t.Reductor().Model()
	if err := t.Render(&buf, "datamatrixlist", &model.Finder, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.String(500, err.Error())
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
