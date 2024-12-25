package create

import (
	_ "embed"
	"fmt"
)

//go:embed vcode.sql
var vCodeDB string

func (vv *versionDb) InsertVCode() (err error) {
	if _, err = vv.db.Exec(vCodeDB); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
