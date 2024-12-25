package a3mssql

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

//go:embed sql/delete_request.sql
var deleteRequestSendedSql string

// старые запросы со статусом Отправлен и Ошибка
func (a *dbAlcohelp3) AdminRequestRemove() (result int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
			result = 0
		}
		a.сloseDB()
	}()

	ctx := context.Background()
	if sqlResult, err := queries.Raw(deleteRequestSendedSql).ExecContext(ctx, a.db); err != nil {
		return sqlResult.RowsAffected()
	} else {
		return sqlResult.RowsAffected()
	}
}
