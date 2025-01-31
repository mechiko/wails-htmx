package stats

import (
	"bytes"
	"firstwails/utility"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
)

func (t *page) upload(c echo.Context) (errOut error) {
	var bufFile bytes.Buffer
	var bufRender bytes.Buffer

	model := t.Reductor().Model()
	model.Stats.Errors = nil
	// model.Stats.CisIn = nil
	model.Stats.File = ""
	model.Stats.CisStatus = make(map[string]int)
	model.Stats.State = 0
	if file, err := c.FormFile("file"); err != nil {
		errOut = fmt.Errorf("ошибка файла %s", err.Error())
	} else {
		model.Stats.File = file.Filename
		src, err := file.Open()
		if err != nil {
			errOut = fmt.Errorf("ошибка файла %s", err.Error())
		} else {
			defer src.Close()
			if _, err = io.Copy(&bufFile, src); err != nil {
				errOut = fmt.Errorf("ошибка файла %s", err.Error())
			}
		}
	}
	if errOut != nil {
		model.Stats.Errors = append(model.Stats.Errors, errOut.Error())
		t.UpdateModel(model, "stats.upload.error")
		if err := t.Render(&bufRender, "page", &model, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.HTML(200, bufRender.String())
	} else {
		cisIn := model.Stats.CisIn
		cisIn = append(cisIn, utility.ReadTextStringArrayReader(&bufFile)...)
		model.Stats.CisIn = utility.UniqueSliceElements(cisIn)
		t.UpdateModel(model, "stats.upload")
		if err := t.Render(&bufRender, "page", &model, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.HTML(200, bufRender.String())
	}
	return nil
}
