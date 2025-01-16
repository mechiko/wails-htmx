package utility

import (
	"github.com/sqweek/dialog"
)

func DialogOpenFile() string {
	result, err := dialog.File().Filter("csv", "csv").Filter("txt", "txt").Filter("all", "*").Load()
	if err != nil {
		return err.Error()
	}
	return result
}

func DialogSaveFile() string {
	result, err := dialog.File().Filter("Excel", "xlsx").Filter("all", "*").Save()
	if err != nil {
		return err.Error()
	}
	return result
}
