package stats

import (
	"bytes"
	"firstwails/domain"
	"firstwails/usecase"
	"firstwails/utility"
	"firstwails/webapp/htmxutil"

	"github.com/labstack/echo/v4"
)

func (t *page) Route(e *echo.Echo) error {
	e.GET("/stats/title", t.Title)
	e.GET("/stats/ready", t.Ready)
	e.GET("/stats/modal", t.modal)
	e.GET("/stats/file", t.file)
	e.GET("/stats/search", t.search)
	e.GET("/stats/reset", t.reset)
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

func (t *page) file(c echo.Context) error {
	file := utility.DialogOpenFile()
	t.Logger().Debugf("stats file selected: %s", file)
	if utility.PathOrFileExists(file) {
		var buf bytes.Buffer
		// обновляем модель Stats
		model := t.Reductor().Model()
		model.Stats.File = file
		msg := domain.Message{}
		msg.Cmd = "model"
		msg.Sender = "stats.router.file"
		msg.Model = &model
		t.Reductor().ChanIn() <- msg
		if err := t.Render(&buf, "page", &model, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.HTML(200, buf.String())
	}
	// возвращаем имя файла для отображения на форме
	// c.String(200, file)
	c.NoContent(204)
	return nil
}

func (t *page) search(c echo.Context) error {
	var buf bytes.Buffer
	// обрабатываем поиск
	model := usecase.New(t).TrueClientSearch(t.Reductor().Model())
	t.Logger().Debugf("stats file selected proccess error: %d", len(model.Error))
	msg := domain.Message{}
	msg.Cmd = "model"
	msg.Sender = "stats.router.search"
	msg.Model = &model
	t.Reductor().ChanIn() <- msg
	// c.String(200, file)
	// c.NoContent(204)
	if err := t.Render(&buf, "page", &model, c); err != nil {
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
	model.Stats.File = ""
	model.Stats.CisList = make(domain.CisSlice, 0)
	model.Stats.Errors = make([]string, 0)
	msg := domain.Message{}
	msg.Cmd = "model"
	msg.Sender = "stats.router.reset"
	msg.Model = &model
	t.Reductor().ChanIn() <- msg
	if err := t.Render(&buf, "page", &model, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
