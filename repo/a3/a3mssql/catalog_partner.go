package a3mssql

import (
	"context"
	"fmt"

	"firstwails/domain"
	"firstwails/repo/a3/a3mssql/a3boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *dbAlcohelp3) CatalogPartner() (list domain.PartnersOriginSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	list = make(domain.PartnersOriginSlice, 0)
	where := `ID in (select distinct max(ID) from partners_egais group by client_reg_id)`
	if sl, err := a3boil.PartnersEgaiss(qm.Where(where)).All(ctx, a.db); err != nil {
		return list, fmt.Errorf("%s %w", modError, err)
	} else {
		for i := range sl {
			// partner := &domain.PartnersOrigin{}
			partner := new(domain.PartnersOrigin)
			partner.ConvertFromA3Mssql(sl[i])
			list = append(list, partner)
		}
		return list, nil
	}
}
