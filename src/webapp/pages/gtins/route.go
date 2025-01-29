package gtins

import (
	"bytes"
	"firstwails/utility"
	"firstwails/webapp/htmxutil"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (t *page) Route(e *echo.Echo) error {
	e.GET("/gtins/title", t.Title)
	e.GET("/gtins/ready", t.Ready)
	e.GET("/gtins/modal", t.modal)
	e.GET("/gtins/file", t.file)
	e.GET("/gtins/search", t.search)
	e.GET("/gtins/reset", t.reset)
	e.GET("/gtins/progress", t.progress)
	e.GET("/gtins/excel/:status", t.excel)
	e.POST("/gtins/upload", t.upload)
	e.POST("/gtins/chunk", t.chunk)
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

func (t *page) file(c echo.Context) error {
	file := utility.DialogOpenFile()
	t.Logger().Debugf("stats file selected: %s", file)
	if utility.PathOrFileExists(file) {
		var buf bytes.Buffer
		// обновляем модель Stats
		model := t.Reductor().Model()
		model.Stats.State = 1
		model.Stats.File = file
		t.UpdateModel(model, "stats.file")
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
	model := t.Reductor().Model()
	model.Stats.CisOut = nil
	model.Stats.Errors = nil
	model.Stats.State = 1
	t.UpdateModel(model, "stats.search")
	// model = usecase.New(t).TrueClientSearch(model)
	t.Search(model)
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
	model.Stats.CisIn = nil
	model.Stats.CisOut = nil
	model.Stats.Errors = nil
	model.Stats.State = 0
	t.UpdateModel(model, "stats.reset")
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

func (t *page) excel(c echo.Context) error {
	status := c.Param("status")
	file := utility.DialogSaveFile()
	file = fmt.Sprintf("%s_%s", file, status)
	allcis := t.Reductor().Model().Stats.CisOut
	arrStatus := make([]string, 0)
	for _, cis := range allcis {
		if cis.Status == status {
			arrStatus = append(arrStatus, cis.Cis)
		}
	}
	size := t.Reductor().Model().Stats.ExcelChunkSize
	if size <= 0 {
		size = 30000
	}
	if err := t.ToExcel(arrStatus, file, size); err != nil {
		return c.String(200, err.Error())
	}
	t.Logger().Debugf("stats file selected: %s", file)
	// возвращаем имя файла для отображения на форме
	out := fmt.Sprintf("записано %d кодов маркировки", len(arrStatus))
	c.String(200, out)
	// c.NoContent(204)
	return nil
}

func (t *page) chunk(c echo.Context) error {
	chunkSize := c.FormValue("chunk")
	model := t.Reductor().Model()
	if size, err := strconv.Atoi(chunkSize); err != nil {
		model.Stats.ExcelChunkSize = 30000
	} else {
		model.Stats.ExcelChunkSize = size
	}
	t.UpdateModel(model, "stats.chunk")
	var buf bytes.Buffer
	if err := t.Render(&buf, "chunk", &model, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, buf.String())
	return nil
}
