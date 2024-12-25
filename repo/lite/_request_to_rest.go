package lite

import (
	"database/sql"
	_ "embed"
	"fmt"
)

//go:embed sql/request_to_rest.sql
var RequestT0RestSql string

func (r *dbLte) RequestToRest(dba3 string, idRequest int64, date string) (err error) {
	defer func() {
		r.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("repo:report RequestToRest panic %v", r)
		}
	}()
	if _, err := r.db.Exec(RequestT0RestSql, sql.Named("dba3", dba3), sql.Named("date1", date), sql.Named("date2", date), sql.Named("request", idRequest)); err != nil {
		return fmt.Errorf("repo:report RequestToRest %w", err)
	}
	return nil
}
