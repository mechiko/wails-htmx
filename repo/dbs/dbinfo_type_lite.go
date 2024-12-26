package dbs

import (
	"fmt"
	"path/filepath"

	"firstwails/domain"
	"firstwails/repo/dbs/sqlite"
)

const defaultLiteDbName = "lite.db"

// база config.db всегда sqlite3 пока

func NewDbInfoLite(cfg *domain.Configuration, config string, scan bool) (di *dbInfo) {
	defer func() {
		if r := recover(); r != nil {
			di.file = ""
			panic(r)
		}
	}()
	di = &dbInfo{
		connection: "",
		mode:       domain.RwModeWithCreate,
		checkMode:  domain.Version,
		driver:     "sqlite",
		file:       cfg.Alcogo4.File,
		name:       cfg.Alcogo4.DbName,
	}
	// пытаемся найти в конфиге приложения или берем по умолчанию
	if di.file == "" {
		di.file = defaultLiteDbName
		di.file = filepath.Join(domain.DbPath, di.file)
	}
	// только для БД котрые с записью создаем так, нужен режим sqlite.RwModeWithCreate
	dbSvc, err := sqlite.New(di)
	if err != nil {
		panic(fmt.Errorf("%s %w", modError, err))
	}
	di.dbService = dbSvc
	return di
}
