package ucexcel

import (
	"firstwails/domain"
	"firstwails/usecase/ucexcel/address"

	"github.com/xuri/excelize/v2"
)

type ucexcel struct {
	domain.IApp
	layout             string
	template           string
	name               string
	address            domain.ExcelAddress
	sheet              string
	file               *excelize.File
	celStyleDefault    int
	celStyleLeft       int
	celStyleVCenter    int
	celStyleBold       int
	celStyleBoldGreen  int
	celStyleBoldBlue   int
	celStyleBoldRight  int
	celStyleBoldCenter int
	celStyle           excelize.Style
}

var _ domain.UseCaseExcel = (*ucexcel)(nil)

func New(app domain.IApp, layout string, template string, name string) domain.UseCaseExcel {
	if layout == "" {
		layout = app.Configuration().Layouts.TimeLayoutDay
	}
	excel := &ucexcel{
		IApp:     app,
		layout:   layout,
		template: template,
		name:     name,
		address:  address.New(14, 0),
	}
	excel.celStyle = excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold:   false,
			Family: "Arial",
			Color:  "000000",
			Size:   9,
		},
	}

	return excel
}
