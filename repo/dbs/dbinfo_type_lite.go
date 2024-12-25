package dbs

import (
	"fmt"
	"path/filepath"

	"firstwails/domain"
	"firstwails/repo/dbs/sqlite"
)

const defaultLiteDbName = "alcogo4lite.db"

// база config.db всегда sqlite3 пока

func NewDbInfoLite(cfg *domain.Configuration, config string, scan bool) (di *dbInfo) {
	defer func() {
		if r := recover(); r != nil {
			di.file = ""
			panic(r)
		}
	}()
	di = &dbInfo{
		driver: cfg.Alcogo4.Driver,
		uri:    sqlite.RwModeWithCreate.String(),
		file:   cfg.Alcogo4.File,
		name:   cfg.Alcogo4.DbName,
	}
	// пытаемся найти в конфиге приложения или берем по умолчанию
	if di.file == "" {
		di.file = defaultLiteDbName
		// di.file = fmt.Sprintf("%s/%s", domain.DbPath, di.file)
		di.file = filepath.Join(domain.DbPath, di.file)
	}
	if di.driver == "" {
		di.driver = "sqlite"
	}
	if di.driver != "sqlite" {
		di.driver = "sqlite"
	}
	if di.uri == "" {
		di.uri = sqlite.RwModeWithCreate.String()
	}
	// только для БД котрые с записью создаем так, нужен режим sqlite.RwModeWithCreate
	dbSvc, err := sqlite.NewDbService(di.file, sqlite.RwModeWithCreate, sqlite.Version)
	if err != nil {
		panic(fmt.Errorf("%s %w", modError, err))
	}
	di.dbService = dbSvc
	return di
}
