package a3sqlite

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

//go:embed sql/clear_rest_volume.sql
var adminClearRestVolumeSql string

// чистим дубли
func (a *dbAlcohelp3) AdminClearRestVolume() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()

	ctx := context.Background()
	if _, err := queries.Raw(adminClearRestVolumeSql).ExecContext(ctx, a.db); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	return nil
}
