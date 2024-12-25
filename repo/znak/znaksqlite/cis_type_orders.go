package znaksqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"firstwails/domain"
	"firstwails/repo/znak/znaksqlite/znakboil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Единица товара
// Групповая потребительская упаковка
// Набор
func (z *dbZnak) CisTypeOrders(cis string) (ois domain.OrderInfoSlice, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic Orders %v", modError, r)
		}
		z.сloseDB()
	}()
	ctx := context.Background()
	ois = make(domain.OrderInfoSlice, 0)
	if orders, err := znakboil.OrderMarkCodes(qm.Where("cis_type = ? and archive <> 1", cis), qm.OrderBy("id desc")).All(ctx, z.db); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return ois, fmt.Errorf("%s CisTypeOrders %w", modError, err)
		}
	} else {
		for i, order := range orders {
			newOrder := &domain.OrderInfo{
				OrderMarkCode: &domain.OrderMarkCode{},
				Guide:         &domain.ProductGuide{},
			}
			newOrder.OrderMarkCode.ConvertFromZnakSqlite(orders[i])
			if guide, err := znakboil.ProductGuides(qm.Where("product_gtin = ?", order.Gtin.String)).One(ctx, z.db); err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					z.Logger().Errorf("%s CisTypeOrders %s", modError, err.Error())
				}
			} else {
				newOrder.Guide.ConvertFromZnakSqlite(guide)
			}
			ois = append(ois, newOrder)
		}
	}
	return ois, nil
}
