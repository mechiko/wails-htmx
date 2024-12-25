package znakmssql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"firstwails/repo/znak/znakmssql/znakboil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Единица товара
// Групповая потребительская упаковка
// Набор
func (z *dbZnak) OrderSerials(id int64) (out []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic Orders %v", modError, r)
		}
		z.сloseDB()
	}()
	out = make([]string, 0)
	ctx := context.Background()
	if serials, err := znakboil.OrderMarkCodesSerialNumbers(qm.Where("id_order_mark_codes = ?", id), qm.OrderBy("id desc")).All(ctx, z.db); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return out, fmt.Errorf("%s OrderSerials %w", modError, err)
		}
	} else {
		for i := range serials {
			out = append(out, serials[i].Code.String)
		}
	}
	return out, nil
}
