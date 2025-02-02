package znaksqlite

import (
	"context"
	"firstwails/domain"
	"firstwails/utility"
	"fmt"
	"time"
)

func (a *dbZnak) FindCis(cis string) (out *domain.OrderMarkCodesSerialNumber, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	out = &domain.OrderMarkCodesSerialNumber{}
	likeCis := fmt.Sprintf("%s%%", utility.TruncateString(cis, 25))
	sql := `select * from order_mark_codes_serial_numbers where code like ? AND status NOT in ('Отправлено');`
	start := time.Now()
	if err = a.db.GetContext(ctx, out, sql, likeCis); err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}
	// out, err = znakboil.OrderMarkCodesSerialNumbers(qm.Where(where, likeCis)).One(ctx, a.db)
	a.Logger().Debugf("execute sql time since %s", time.Since(start))
	return out, nil
}
