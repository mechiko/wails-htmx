package lite

import (
	_ "embed"
	"fmt"
)

func (r *dbReport) DelEmptyRegIdLicense() (err error) {
	defer func() {
		r.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	query := `DELETE FROM licenses WHERE client_reg_id='';`
	if _, err := r.db.Exec(query); err != nil {
		return fmt.Errorf("dbreport:DelEmptyRegIdLicense %w", err)
	}
	return nil
}
