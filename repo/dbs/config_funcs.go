package dbs

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

// планировал переименовать но решил оставить ...
func (dd *dbs) driverConfig(key string) string {
	driver := strings.ToLower(dd.fromConfig(key))
	if driver == "sqlite" {
		return "sqlite"
	}
	return driver
}

func (dd *dbs) fromConfig(key string) (out string) {
	defer func() {
		if r := recover(); r != nil {
			out = ""
		}
	}()
	params := ""
	if dd.configDb == nil {
		return ""
	}
	if dd.configDb.absent {
		return ""
	}
	dbSql, err := dd.configDb.InfoDbService().Db()
	if err != nil {
		return ""
	}
	defer func() {
		dbSql.Close()
		dd.configDb.InfoDbService().Close()
	}()
	db := sqlx.NewDb(dbSql, "sqlite3")
	sql := "select TRIM(value) as value from parameters where name=?;"
	err = db.Get(&params, sql, key)
	if err != nil {
		return ""
	}
	return params
}
