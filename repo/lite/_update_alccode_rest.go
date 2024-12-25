package lite

import (
	_ "embed"
	"fmt"

	"firstwails/entity"
)

//go:embed sql/update_alccode_rest.sql
var UpdateAlcCodeRestSql string

func (r *dbLte) UpdateAlcCodeRest(rest *entity.ReportAlcCodeRest) (err error) {
	defer func() {
		// обязательно, иначе лочится рессурс репозитория для указателя на db в DbService
		r.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()

	if _, err := r.db.Exec(UpdateAlcCodeRestSql, rest.AlcCode, rest.AlcVolume, rest.DateRest, rest.Volume); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
