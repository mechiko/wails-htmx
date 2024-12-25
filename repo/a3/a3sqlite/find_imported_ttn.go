package a3sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"firstwails/repo/a3/a3sqlite/a3boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ищем любой тип с таким же номером и датой со state <> 'Создан'
// если такие есть то возвращается ID
// ошибка если только реально ошибка обращения к БД
func (a *dbAlcohelp3) FindImportTtn(number, date string) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic FindImportTtn %v", modError, r)
		}
		a.сloseDB()
	}()
	ctx := context.Background()
	where := `doc_number = ? and doc_date = ? and state <> 'Создан'` // and status in ('Подтверждён', 'Проведён')`
	if ttn, err := a3boil.TTNS(qm.Where(where, number, date)).One(ctx, a.db); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("%s FindImportTtn %w", modError, err)
	} else {
		return int(ttn.ID), nil
	}
}
