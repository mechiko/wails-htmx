package dbs

import (
	"fmt"
	"path/filepath"
	"regexp"

	"firstwails/domain"
	"firstwails/repo/dbs/mssql"
	"firstwails/repo/dbs/sqlite"
	"firstwails/utility"
)

func NewDbInfoA3(cfg *domain.Configuration, dbType string) (di *dbInfo) {
	defer func() {
		if r := recover(); r != nil {
			di.file = ""
			di.absent = true
			panic(r)
		}
	}()
	di = &dbInfo{
		connection: "",
		mode:       domain.RwMode, // возможна запись в базу УТМ
		checkMode:  domain.NoMatter,
		driver:     cfg.Alcohelp3.Driver,
		file:       cfg.Alcohelp3.File,
		name:       cfg.Alcohelp3.DbName,
	}
	if di.driver == "" {
		if di.driver = dbType; di.driver == "" {
			di.driver = "sqlite"
		}
	}
	switch di.driver {
	case "mssql":
		di.file = ""
		if di.connection == "" {
			di.connection = cfg.Mssql.Connection
		}
		if di.name == "" {
			if di.name = cfg.Application.Fsrarid; di.name == "" {
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
		di.name = ""
		if di.file == "" {
			if di.file = cfg.Application.Fsrarid; di.file == "" {
				if di.file = di.findA3Name(); di.file == "" {
					di.absent = true
					return di
				} else {
					di.file = di.file + ".db"
				}
			} else {
				// если есть фсрар ид
				di.file = di.file + ".db"
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
	default:
		di.absent = true
	}
	return di
}

func (di *dbInfo) findA3DbName() string {
	re, err := regexp.Compile(`^0[0-9]{11}\.db$`)
	if err != nil {
		return ""
	}
	if files, err := utility.FilteredSearchOfDirectoryTree(re, ""); err != nil {
		return ""
	} else {
		if len(files) == 0 {
			return ""
		}
		return files[0]
	}
}

func (di *dbInfo) findA3Name() string {
	findName := di.findA3DbName()
	if findName == "" {
		return ""
	}
	_, file := filepath.Split(findName)
	before := file[0 : len(file)-len(filepath.Ext(file))]
	// before, _ := strings.CutSuffix(file, filepath.Ext(file))
	return before
}
