package dbs

import (
	"fmt"

	"firstwails/domain"
	"firstwails/repo/dbs/mssql"
	"firstwails/repo/dbs/sqlite"
	"firstwails/utility"
)

func NewDbInfo4z(cfg *domain.Configuration, name string, dbType string) (di *dbInfo) {
	defer func() {
		if r := recover(); r != nil {
			di.file = ""
			panic(r)
		}
	}()
	di = &dbInfo{
		driver: cfg.TrueZnak.Driver,
		uri:    "",
		file:   cfg.TrueZnak.File,
		name:   cfg.TrueZnak.DbName,
	}
	if di.driver == "" {
		if di.driver = dbType; di.driver == "" {
			di.driver = "sqlite"
		}
	}
	switch di.driver {
	case "mssql":
		if di.name == "" {
			if di.name = name; di.name == "" {
				di.absent = true
				return di
			}
		}
		if di.uri == "" {
			di.uri = cfg.Mssql.ConnectionUri
		}
		di.file = ""
		if di.name == "" {
			if di.name = name; di.name == "" {
				di.absent = true
				return di
			}
		}
		dbs, err := mssql.New(di.name, di.uri)
		if err != nil {
			panic(fmt.Errorf("%s %w", modError, err))
		}
		di.dbService = dbs
		return di
	case "sqlite":
		if di.file == "" {
			if di.file = di.name; di.file == "" {
				if di.file = name; di.file == "" {
					di.absent = true
					return di
				} else {
					di.file = fmt.Sprintf("%s.db", di.file)
				}
			} else {
				di.file = fmt.Sprintf("%s.db", di.file)
			}
		}
		if utility.PathOrFileExistsMust(di.file) {
			if di.uri == "" {
				di.uri = sqlite.RoMode.String()
			}
			dbs, err := sqlite.NewFromUri(di.file, di.uri, sqlite.NoMatter)
			if err != nil {
				panic(fmt.Errorf("%s %w", modError, err))
			}
			di.dbService = dbs
		} else {
			di.absent = true
		}
		return di
	default:
		di.absent = true
	}
	return di
}
