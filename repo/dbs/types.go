package dbs

import (
	"strings"

	"firstwails/domain"
)

const modError = "dbs:dbinfo"

type dbs struct {
	liteDb        *dbInfo
	a3Db          *dbInfo
	znakDb        *dbInfo
	configDb      *dbInfo
	scan          bool
	configFile    string
	configuration *domain.Configuration
}

// config файл бд config.db
// scan надо ли сканировать имена файлов по шаблону
func New(cfg *domain.Configuration, config string, scan bool) *dbs {
	d := &dbs{
		configuration: cfg,
		configFile:    config,
		scan:          scan,
	}
	d.configDb = NewDbInfoConfig(cfg)
	file4z := d.fromConfig("oms_id")
	dbType := strings.ToLower(d.fromConfig("db_type"))
	if dbType == "" {
		dbType = cfg.Application.DbType
	}
	d.znakDb = NewDbInfo4z(cfg, file4z, dbType)
	d.a3Db = NewDbInfoA3(cfg, dbType)
	d.liteDb = NewDbInfoLite(cfg, "", scan)
	return d
}

func (d *dbs) Lite() domain.DbInfo {
	return d.liteDb
}

func (d *dbs) Config() domain.DbInfo {
	return d.configDb
}

func (d *dbs) Znak() domain.DbInfo {
	return d.znakDb
}

func (d *dbs) A3() domain.DbInfo {
	return d.a3Db
}

func (d *dbs) ZnakSvc() domain.DbService {
	return d.znakDb.InfoDbService()
}

func (d *dbs) A3Svc() domain.DbService {
	return d.a3Db.InfoDbService()
}

func (d *dbs) LiteSvc() domain.DbService {
	return d.liteDb.InfoDbService()
}

func (d *dbs) ConfigSvc() domain.DbService {
	return d.configDb.InfoDbService()
}

func (d *dbs) SaveConfig(cfg domain.IConfig) (err error) {
	err = d.configDb.SaveConfig(cfg, "config")
	if err != nil {
		return
	}
	err = d.a3Db.SaveConfig(cfg, "alcohelp3")
	if err != nil {
		return
	}
	err = d.znakDb.SaveConfig(cfg, "trueznak")
	if err != nil {
		return
	}
	err = d.liteDb.SaveConfig(cfg, "alcogo4")
	if err != nil {
		return
	}
	return
}
