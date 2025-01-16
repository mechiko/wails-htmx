package stats

import (
	"firstwails/usecase/ucexcel"
	"fmt"
)

func (t *page) ToExcel(ar []string, name string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic gui excel %v", r)
		}
	}()
	excel := ucexcel.New(t, "", "", name)
	if err := excel.Open(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := excel.ReportList(ar); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := excel.SaveSimple(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil

}
