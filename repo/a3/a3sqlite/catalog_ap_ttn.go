package a3sqlite

import (
	"context"
	_ "embed"
	"fmt"

	"firstwails/domain"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

func (a *dbAlcohelp3) CatalogApTTNMap() (out map[string]*domain.ApEgaisImport, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	out = make(map[string]*domain.ApEgaisImport)
	apSlice := make(domain.ApEgaisSlice, 0)
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

func (a *dbAlcohelp3) CatalogApTTNTypeMap() (out map[string]string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	out = make(map[string]string)
	apSlice := make(domain.ApEgaisSlice, 0)
	if err := queries.Raw(apFromTTNSql).Bind(ctx, a.db, &apSlice); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	} else {
		for _, ap := range apSlice {
			if _, ok := out[ap.ProductAlcCode]; !ok {
				out[ap.ProductAlcCode] = ap.ProductUnitType
			}
		}

	}
	return out, nil
}

func (a *dbAlcohelp3) CatalogApTTN() (out domain.ApEgaisSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	out = make(domain.ApEgaisSlice, 0)
	if err := queries.Raw(apFromTTNSql).Bind(ctx, a.db, &out); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	return out, nil
}
