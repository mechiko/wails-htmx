package a3sqlite

import (
	"context"
	"fmt"

	"firstwails/repo/a3/a3sqlite/a3boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// удаляем по условию doc_number = ? and doc_date =? and state = 'Создан' and status = 'Не проведён' and ttn_type='Импорт'
func (a *dbAlcohelp3) RemoveTTN(ttn string, date string) (err error) {
	ctx := context.Background()
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("panic %v", r)
		}
		a.сloseDB()
	}()

	// ищем все ТТН по условию
	if ttnFind, err := a3boil.TTNS(qm.Where("doc_number = ? and doc_date =? and state = 'Создан' and status = 'Не проведён' and ttn_type='Импорт'", ttn, date)).All(ctx, tx); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		for _, ttn := range ttnFind {
			if _, err := a3boil.TTNProducts(qm.Where("id_ttn = ?", ttn.ID)).DeleteAll(ctx, tx); err != nil {
				return fmt.Errorf("%w", err)
			}
			if _, err := ttn.Delete(ctx, tx); err != nil {
				tx.Rollback()
				return fmt.Errorf("%w", err)
			}
		}
		tx.Commit()
	}
	return nil
}
