package znakmssql

import (
	"context"
	"firstwails/repo/znak/znakmssql/znakboil"
	"fmt"
)

func (a *dbZnak) GtinAll() (out []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	if sl, err := znakboil.ProductGuides().All(ctx, a.db); err != nil {
		return nil, fmt.Errorf("%w", err)
	} else {
		out = make([]string, 0)
		for i := range sl {
			out = append(out, sl[i].ProductGtin.String)
		}
		return out, nil
	}
}
