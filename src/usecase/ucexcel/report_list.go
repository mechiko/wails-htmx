package ucexcel

import "firstwails/usecase/ucexcel/address"

// var countRow int

func (ue *ucexcel) ReportList(report []string) error {
	// https://xuri.me/excelize/ru/workbook.html#GetSheetProps
	ue.sheet = "Sheet1"
	ue.address = address.New(1, 0)
	for _, ss := range report {
		ue.templateReportListLine(ss)
	}

	return nil
}

func (ue *ucexcel) templateReportListLine(s string) {
	addr := ue.address.Address()
	if err := ue.file.SetCellStr(ue.sheet, addr, s); err != nil {
		ue.Logger().Errorf("excel error %s", err.Error())
	}
	ue.address.NextRow()
}
