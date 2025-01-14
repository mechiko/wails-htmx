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
		connection: "",
		mode:       domain.RoMode,
		checkMode:  domain.NoMatter,
		driver:     cfg.TrueZnak.Driver,
		file:       cfg.TrueZnak.File,
		name:       cfg.TrueZnak.DbName,
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
		if di.connection == "" {
			di.connection = cfg.Mssql.Connection
		}
		di.file = ""
		if di.name == "" {
			if di.name = name; di.name == "" {
				di.absent = true
				return di
			}
		}
		dbs, err := mssql.New(di)
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
			dbs, err := sqlite.New(di)
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
