package a3sqlite

import (
	"context"
	"fmt"

	"firstwails/repo/a3/a3sqlite/a3boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *dbAlcohelp3) CatalogApTypeMap() (mapType map[string]string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	mapType = make(map[string]string)
	ctx := context.Background()
	where := `ID in (select distinct max(ID) from ap_egais group by product_alc_code)`
	if sl, err := a3boil.ApEgaiss(qm.Where(where)).All(ctx, a.db); err != nil {
		return mapType, fmt.Errorf("%s %w", modError, err)
	} else {
		for _, ap := range sl {
			if _, ok := mapType[ap.ProductAlcCode.String]; !ok {
				mapType[ap.ProductAlcCode.String] = ap.ProductUnitType.String
			}
		}
	}
	return mapType, nil
}
