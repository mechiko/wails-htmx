package znakmssql

import (
	"fmt"
)

func (a *dbZnak) AttachLite(liteDbName string, status string, gtin string) (out int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	// out = make([]string, 0)
	// if liteDbName == "" {
	// 	return out, fmt.Errorf("%s укажите имя бд lite", modError)
	// }
	// if !utility.PathOrFileExists(liteDbName) {
	// 	return out, fmt.Errorf("%s бд lite не найдена", modError)
	// }
	// attachSql := fmt.Sprintf("ATTACH DATABASE %s AS litedb;", liteDbName)
	// sqlLite := `select cr.cis from litedb.cis_request cr where cr.status = ?;`
	// sql := fmt.Sprintf("%s\n%s", attachSql, sqlLite)
	// err = a.db.Select(&out, sql, status)
	// if err != nil {
	// 	return out, fmt.Errorf("%s %w", modError, err)
	// }

	return 0, nil
}
