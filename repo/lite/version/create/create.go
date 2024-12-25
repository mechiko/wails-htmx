package create

import (
	_ "embed"
	"fmt"
)

const createVersion string = `
CREATE TABLE if not exists dboptions (
	name TEXT NOT NULL DEFAULT '',
	value TEXT NOT NULL DEFAULT '',
	PRIMARY KEY (name)
);`

const setVersion string = `INSERT OR REPLACE INTO dboptions (name, value) VALUES ('version',?)`

const getVersion string = `select value from dboptions where name = 'version' limit 1;`

//go:embed createDB.sql
var createDB string

//go:embed index.sql
var indexDB string

//go:embed insert.sql
var insertDB string

func (vv *versionDb) create() error {
	if err := vv.createScheme(); err != nil {
		return fmt.Errorf("repo:create %w", err)
	}
	if err := vv.createData(); err != nil {
		return fmt.Errorf("repo:create %w", err)
	}
	return nil
}

// create() создаем БД c нуля
func (vv *versionDb) createScheme() (err error) {
	if _, err = vv.db.Exec(createVersion); err != nil {
		return fmt.Errorf("%w", err)
	}
	if _, err = vv.db.Exec(createDB); err != nil {
		return fmt.Errorf("%w", err)
	}
	if _, err = vv.db.Exec(indexDB); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (vv *versionDb) createData() (err error) {
	if _, err = vv.db.Exec(insertDB); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := vv.InsertVCode(); err != nil {
		return fmt.Errorf("repo:upgrade %w", err)
	}
	if err = vv.InsertCountry(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
