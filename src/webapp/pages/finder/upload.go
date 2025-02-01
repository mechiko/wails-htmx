package finder

import (
	"bytes"
	"firstwails/domain"
	"firstwails/utility"
	"fmt"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
)

func (t *page) upload(c echo.Context) (errOut error) {
	var bufFile bytes.Buffer
	var bufRender bytes.Buffer
	model := t.Reductor().Model()
	model.Finder.Errors = nil
	model.Finder.File = ""
	model.Finder.State = 0
	if file, err := c.FormFile("file"); err != nil {
		errOut = fmt.Errorf("ошибка файла %s", err.Error())
	} else {
		model.Finder.File = file.Filename
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
		model.Finder.Errors = append(model.Finder.Errors, errOut.Error())
		t.UpdateModel(model, "gtins.upload.error")
		if err := t.Render(&bufRender, "page", &model, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.HTML(200, bufRender.String())
	} else {
		arrayFromFile := utility.ReadTextStringArrayReader(&bufFile)
		for i, line := range arrayFromFile {
			if line == "" {
				continue
			}
			s := strings.Split(line, " ")
			if len(s) != 4 {
				err := fmt.Sprintf("%s строка %d содержит неправильный формат должно быть 4 поля [%s]", modError, i, line)
				t.Logger().Errorf("%s %s", modError, err)
				model.Finder.Errors = append(model.Finder.Errors, err)
				continue
			}
			info := domain.CisFindInfo{
				Root:   s[0],
				Dir:    s[1],
				Name:   s[2],
				CisSrc: s[3],
				Cis:    utility.TruncateString(s[3], 25),
			}
			if len(info.Cis) < 25 {
				err := fmt.Sprintf("%s строка %d [%s] поиск ошибка КМ меньше 25", modError, i, line)
				t.Logger().Errorf("%s %s", modError, err)
			}
			if serial, err := t.Repo().DbZnak().FindCis(info.Cis); err != nil {
				err := fmt.Sprintf("%s строка %d [%s] поиск ошибка %s", modError, i, line, err.Error())
				t.Logger().Errorf("%s %s", modError, err)
				model.Finder.Errors = append(model.Finder.Errors, err)
			} else {
				info.Code = serial.Code
				info.Order = serial.IDOrderMarkCodes
				// var b byte = 232
				// b = 235 + (232-128)
				// info.CodeFNS = "\f"
				info.CodeFNS += serial.Code
				// t.Logger().Debugf("FNS [%s] %d", info.CodeFNS, len(info.CodeFNS))
			}
			model.Finder.CisFindInfoIn = append(model.Finder.CisFindInfoIn, &info)
		}
		t.UpdateModel(model, "finder.upload")
		if err := t.Render(&bufRender, "page", &model.Finder, c); err != nil {
			t.Logger().Errorf("%s %s", modError, err.Error())
			c.NoContent(204)
			return nil
		}
		c.HTML(200, bufRender.String())
	}
	return nil
}
