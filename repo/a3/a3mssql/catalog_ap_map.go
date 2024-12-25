package a3mssql

import (
	"context"
	"fmt"

	"firstwails/domain"
	"firstwails/repo/a3/a3mssql/a3boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *dbAlcohelp3) CatalogApMap() (mapType map[string]*domain.ApEgaisImport, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	mapType = map[string]*domain.ApEgaisImport{}
	ctx := context.Background()
	where := `ID in (select distinct max(ID) from ap_egais group by product_alc_code)`
	if sl, err := a3boil.ApEgaiss(qm.Where(where)).All(ctx, a.db); err != nil {
		return mapType, fmt.Errorf("%s %w", modError, err)
	} else {
		for _, ap := range sl {
			if _, ok := mapType[ap.ProductAlcCode.String]; !ok {
				apEgais := domain.ApEgaisImport{}
				apEgais.ConvertFromA3Mssql(ap)
				mapType[ap.ProductAlcCode.String] = &apEgais
			}
		}
	}
	return mapType, nil
}
