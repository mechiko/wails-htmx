package v002

import (
	_ "embed"
	"fmt"

	// этот драйвер не зависит от CGO поэтому вроде удобно
	// но проблема для 32 бит
	// _ "modernc.org/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

const setVersion string = `INSERT OR REPLACE INTO dboptions (name, value) VALUES ('version',?)`

const getVersion string = `select value from dboptions where name = 'version' limit 1;`

//go:embed createDB.sql
var createDB string

//go:embed index.sql
var indexDB string

//go:embed drop.sql
var dropDB string

func (vv *versionDb) create() (err error) {
	if err := vv.createScheme(); err != nil {
		return fmt.Errorf("repo:v002 %w", err)
	}
	if err := vv.createData(); err != nil {
		return fmt.Errorf("repo:v002 %w", err)
	}
	return nil
}

// create() создаем БД c нуля
func (vv *versionDb) createScheme() (err error) {
	if _, err = vv.db.Exec(dropDB); err != nil {
		return fmt.Errorf("%w", err)
	}
	if _, err = vv.db.Exec(createDB); err != nil {
		return fmt.Errorf("%w", err)
	}
	if _, err = vv.db.Exec(indexDB); err != nil {
		return fmt.Errorf("%w", err)
	}
	return err
}

func (vv *versionDb) createData() error {
	return nil
}
