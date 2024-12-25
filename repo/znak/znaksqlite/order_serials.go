package znaksqlite

import (
	"database/sql"
	"errors"
	"fmt"
)

// Единица товара
// Групповая потребительская упаковка
// Набор
func (z *dbZnak) OrderSerials(id int64) (out []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic Orders %v", modError, r)
		}
		z.сloseDB()
	}()
	out = make([]string, 0)
	sqlOrder := `select code 
from order_mark_codes_serial_numbers omcsn 
where id_order_mark_codes = ?;`
	if err := z.db.Select(&out, sqlOrder, id); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return out, fmt.Errorf("%s %w", modError, err)
		}
	}
	return out, nil
}
