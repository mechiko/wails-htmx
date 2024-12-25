package a3mssql

import (
	"context"
	"fmt"

	"firstwails/domain"
	"firstwails/repo/a3/a3mssql/a3boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *dbAlcohelp3) CatalogAp() (list domain.ApEgaisSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	list = make(domain.ApEgaisSlice, 0)
	where := `ID in (select distinct max(ID) from dbo.ap_egais group by product_alc_code)`
	if sl, err := a3boil.ApEgaiss(qm.Where(where)).All(ctx, a.db); err != nil {
		return list, fmt.Errorf("%s %w", modError, err)
	} else {
		for i := range sl {
			item := &domain.ApEgaisImport{}
			item.ConvertFromA3Mssql(sl[i])
			list = append(list, item)
		}
		return list, nil
	}
}
