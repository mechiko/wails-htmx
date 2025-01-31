package gtins

import (
	"bytes"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
)

func (t *page) upload(c echo.Context) (errOut error) {
	var buf bytes.Buffer

	model := t.Reductor().Model()
	model.Gtins.Errors = nil
	// model.Gtins.CisIn = nil
	model.Gtins.File = ""
	// model.Gtins.CisStatus = make(map[string]int)
	model.Gtins.State = 0
	if file, err := c.FormFile("file"); err != nil {
		errOut = fmt.Errorf("ошибка файла %s", err.Error())
	} else {
		model.Gtins.File = file.Filename
		src, err := file.Open()
		if err != nil {
			errOut = fmt.Errorf("ошибка файла %s", err.Error())
		} else {
			defer src.Close()
			if _, err = io.Copy(&buf, src); err != nil {
				errOut = fmt.Errorf("ошибка файла %s", err.Error())
			}
		}
	}
	if errOut != nil {
		model.Gtins.Errors = append(model.Gtins.Errors, errOut.Error())
		t.UpdateModel(model, "gtins.upload.error")
		if err := t.Render(&buf, "page", &model, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.HTML(200, buf.String())
	} else {
		t.UpdateModel(model, "gtins.upload")
		if err := t.Render(&buf, "page", &model, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.HTML(200, buf.String())
	}
	return nil
}
