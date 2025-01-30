package lite

import "fmt"

func (a *dbLite) RequestAttach() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()

	sql := ""
	if _, err = a.db.Exec(sql); err != nil {
		return fmt.Errorf("%s %v", modError, err)
	}
	return nil
}
