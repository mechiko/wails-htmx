package lite

import (
	_ "embed"
	"fmt"

	"firstwails/entity"
)

//go:embed sql/get_all_alccode_rest.sql
var GetAllAlcCodeRestSql string

func (r *dbReport) GetAllAlcCodeRestMap(start, end string) (mapRest entity.ReportAlcCodeRestMap, err error) {
	defer func() {
		// обязательно, иначе лочится рессурс репозитория для указателя на db в DbService
		r.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()

	mapRest = make(entity.ReportAlcCodeRestMap)
	arr := []entity.ReportAlcCodeRest{}
	err = r.db.Select(&arr, GetAllAlcCodeRestSql, start, end)
	if err != nil {
		return mapRest, fmt.Errorf("%w", err)
	}
	if len(arr) > 0 {
		for _, v := range arr {
			if v.AlcCode == "" {
				return mapRest, fmt.Errorf("report:restmap alc_code empty")
			}
			if v.DateRest == "" {
				return mapRest, fmt.Errorf("report:restmap date_rest empty")
			}
			key := v.AlcCode + ":" + v.AlcVolume + ":" + v.DateRest
			mapRest[key] = v.Volume
		}
	}

	return mapRest, nil
}
