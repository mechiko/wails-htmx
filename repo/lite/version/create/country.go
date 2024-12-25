package create

import (
	_ "embed"
)

//go:embed country.sql
var countryDB string

func (vv *versionDb) InsertCountry() (err error) {
	_, err = vv.db.Exec(countryDB)
	return err
}
