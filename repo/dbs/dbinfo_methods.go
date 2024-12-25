package dbs

import (
	"fmt"

	"firstwails/domain"
)

func (di *dbInfo) Name() string {
	return di.name
}

func (di *dbInfo) File() string {
	return di.file
}

func (di *dbInfo) Driver() string {
	return di.driver
}

func (di *dbInfo) Uri() string {
	return di.uri
}

func (di *dbInfo) Absent() bool {
	return di.absent
}

func (di *dbInfo) InfoDbService() domain.DbService {
	if di.dbService == nil {
		panic(fmt.Errorf("%s DbService is nil uri=%s", modError, di.uri))
	}
	return di.dbService
}

func (di *dbInfo) SaveConfig(cfg domain.IConfig, key string) (err error) {
	// err = cfg.Set(key+".connectionuri", di.uri, true)
	// if err != nil {
	// 	return
	// }
	err = cfg.Set(key+".dbname", di.name, true)
	if err != nil {
		return
	}
	// err = cfg.Set(key+".driver", di.driver, true)
	// if err != nil {
	// 	return
	// }
	err = cfg.Set(key+".file", di.file, true)
	if err != nil {
		return
	}
	return
}

func (di *dbInfo) String() string {
	if di.absent {
		return fmt.Sprintf("Driver:%s DB:отсутствует", di.driver)
	}
	switch di.driver {
	case "mssql":
		return fmt.Sprintf("Driver:mssql DB:%s Server:%s", di.name, di.uri)
	case "sqlite":
		return fmt.Sprintf("Driver:sqlite DB:%s File:%s", di.name, di.file)
	default:
		return ""
	}
}
