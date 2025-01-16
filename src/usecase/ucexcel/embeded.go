package ucexcel

import (
	_ "embed"
)

//go:embed templates/utilizedReport.xlsx
var utilizedReportExcel []byte

//go:embed templates/utilizedReportGtin.xlsx
var utilizedReportGtinExcel []byte
