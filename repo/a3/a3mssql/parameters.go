package a3mssql

import "fmt"

func (a *dbAlcohelp3) Parameters() (s []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
		a.сloseDB()
	}()
	s = make([]string, 0)
	db := a.db
	rows, err := db.Query("SELECT * FROM dbo.parameters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return s, err
		}
		s = append(s, version)
	}
	if err = rows.Err(); err != nil {
		return s, err
	}
	return s, nil
}
