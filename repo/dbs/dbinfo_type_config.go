package dbs

import (
	"fmt"

	"firstwails/domain"
	"firstwails/repo/dbs/sqlite"
	"firstwails/utility"
)

const defaultConfigDbName = "config.db"

// база config.db всегда sqlite3 пока

func NewDbInfoConfig(cfg *domain.Configuration) (di *dbInfo) {
	defer func() {
		if r := recover(); r != nil {
			di.file = ""
			panic(r)
		}
	}()
	di = &dbInfo{
		driver: cfg.Config.Driver,
		uri:    cfg.Sqlite.ConnectionUri,
		file:   cfg.Config.File,
		name:   cfg.Config.DbName,
	}
	// пытаемся найти в конфиге приложения или берем по умолчанию
	if di.file == "" {
		di.file = defaultConfigDbName
	}
	if utility.PathOrFileExistsMust(di.file) {
		if di.uri == "" {
			di.uri = cfg.Sqlite.ConnectionUri
		}
		dbs, err := sqlite.NewFromUri(di.file, di.uri, sqlite.NoMatter)
		if err != nil {
			panic(fmt.Errorf("%s %w", modError, err))
		}
		di.dbService = dbs
	} else {
		// если файла нет физически
		di.absent = true
	}
	return di
}
