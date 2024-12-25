package a3sqlite

import (
	"context"
	_ "embed"
	"fmt"

	"firstwails/domain"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

//go:embed sql/ap_from_ttn.sql
var apFromTTNSql string

//go:embed sql/ap_from_rest.sql
var apFromRestSql string

//go:embed sql/ap_from_ap.sql
var apFromApSql string

func (a *dbAlcohelp3) CatalogApUnionMap() (out map[string]*domain.ApEgaisImport, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	out = make(map[string]*domain.ApEgaisImport)
	apSlice := make(domain.ApEgaisSlice, 0)
	if err := queries.Raw(apFromApSql).Bind(ctx, a.db, &apSlice); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	} else {
		for i, ap := range apSlice {
			if _, ok := out[ap.ProductAlcCode]; !ok {
				out[ap.ProductAlcCode] = apSlice[i]
			}
		}

	}
	apSlice = nil
	if err := queries.Raw(apFromRestSql).Bind(ctx, a.db, &apSlice); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	} else {
		for i, ap := range apSlice {
			if _, ok := out[ap.ProductAlcCode]; !ok {
				out[ap.ProductAlcCode] = apSlice[i]
			}
		}

	}
	apSlice = nil
	if err := queries.Raw(apFromTTNSql).Bind(ctx, a.db, &apSlice); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	} else {
		for i, ap := range apSlice {
			if _, ok := out[ap.ProductAlcCode]; !ok {
				out[ap.ProductAlcCode] = apSlice[i]
			}
		}

	}
	return out, nil
}
