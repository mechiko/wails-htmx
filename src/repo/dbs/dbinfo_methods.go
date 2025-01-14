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

func (di *dbInfo) Absent() bool {
	return di.absent
}

func (di *dbInfo) InfoDbService() domain.IDbService {
	if di.dbService == nil {
		panic(fmt.Errorf("%s DbService is nil uri=%s", modError, di.URI()))
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
		return fmt.Sprintf("Driver:mssql DB:%s Server:%s", di.name, di.URI())
	case "sqlite":
		return fmt.Sprintf("Driver:sqlite File %s URI:%s", di.file, di.URI())
	default:
		return ""
	}
}

func (di *dbInfo) Mode() domain.UriMode {
	return di.mode
}

func (di *dbInfo) CheckMode() domain.CheckMode {
	return di.checkMode
}

func (di *dbInfo) URI() string {
	uri := ""
	switch di.driver {
	case "sqlite":
		uri = fmt.Sprintf("file:%s?%s", di.file, di.mode)
	case "mssql":
		uri = fmt.Sprintf("%s;database=%s", di.connection, di.name)
	}
	return uri
}
