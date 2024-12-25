package lite

import (
	_ "embed"
	"fmt"

	"firstwails/entity"
)

func (r *) DeleteLicense(l *entity.License) (err error) {
	defer func() {
		// обязательно, иначе лочится рессурс репозитория для указателя на db в DbService
		r.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()

	query := `DELETE FROM licenses WHERE client_reg_id = ?;`
	if _, err := r.db.Exec(query, l.ClientRegID); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
