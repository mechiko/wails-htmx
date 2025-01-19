package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (c *dbConfig) Key(key string) (out string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		c.сloseDB()
	}()
	params := ""
	db := sqlx.NewDb(c.db, "sqlite3")
	sql := "select TRIM(value) as value from parameters where name=?;"
	err = db.Get(&params, sql, key)
	if err != nil {
		return "", nil
	}
	return params, nil
}
