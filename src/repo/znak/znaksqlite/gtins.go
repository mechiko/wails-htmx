package znaksqlite

import (
	"fmt"
)

func (a *dbZnak) GtinAll() (out []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	// ctx := context.Background()

	out = make([]string, 0)

	sql := `select DISTINCT product_gtin from product_guides;`
	err = a.db.Select(&out, sql)
	if err != nil {
		return out, fmt.Errorf("%s %w", modError, err)
	}

	return out, nil
}
