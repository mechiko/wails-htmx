package finder

import (
	"bytes"
	"firstwails/domain"
	"firstwails/utility"
	"fmt"

	"github.com/labstack/echo/v4"
)

func (t *page) find(c echo.Context) (errOut error) {
	var bufRender bytes.Buffer
	cisField := c.FormValue("cisfield")
	model := t.Reductor().Model()
	model.Finder.FormInput = cisField
	model.Finder.Errors = nil
	model.Finder.File = ""
	model.Finder.State = 0
	// model.Finder.CisFindInfoIn = nil
	info := domain.CisFindInfo{
		CisSrc: cisField,
		Cis:    utility.TruncateString(cisField, 25),
	}
	// if len(info.Cis) < 25 {
	// 	err := fmt.Sprintf("%s строка поиск КМ меньше 25", modError, i, line)
	// 	t.Logger().Errorf("%s %s", modError, err)
	// 	info.Code = "строка меньше 25"
	// }
	if serial, err := t.Repo().DbZnak().FindCis(info.Cis); err != nil {
		errStr := fmt.Sprintf("%s поиск ошибка %s", modError, err.Error())
		t.Logger().Errorf("%s %s", modError, errStr)
		model.Finder.Errors = append(model.Finder.Errors, errStr)
		info.Code = err.Error()
	} else {
		info.Code = serial.Code
		info.Order = serial.IDOrderMarkCodes
		info.CodeFNS += serial.Code
	}
	model.Finder.CisFindInfoIn = append(model.Finder.CisFindInfoIn, &info)
	t.UpdateModel(model, "finder.find")
	if err := t.Render(&bufRender, "page", &model.Finder, c); err != nil {
		t.Logger().Errorf("%s %s", modError, err.Error())
		c.NoContent(204)
		return nil
	}
	c.HTML(200, bufRender.String())
	return nil
}
