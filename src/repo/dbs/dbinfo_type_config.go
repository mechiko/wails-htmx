package dbs

import (
	"fmt"

	"firstwails/domain"
	"firstwails/repo/dbs/sqlite"
	"firstwails/utility"
)

// база config.db всегда sqlite3 пока

func NewDbInfoConfig(cfg *domain.Configuration) (di *dbInfo) {
	defer func() {
		if r := recover(); r != nil {
			di.file = ""
			di.absent = true
			panic(r)
		}
	}()
	di = &dbInfo{
		driver:     cfg.Config.Driver,
		connection: "",
		file:       cfg.Config.File,
		name:       cfg.Config.DbName,
		mode:       domain.RoMode,
		checkMode:  domain.NoMatter,
	}
	// пытаемся найти в конфиге приложения или берем по умолчанию
	if di.file == "" {
		di.file = defaultConfigDbName
	}
	if utility.PathOrFileExistsMust(di.file) {
		dbs, err := sqlite.New(di)
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
