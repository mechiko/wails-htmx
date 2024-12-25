package domain

import "database/sql"

type DbService interface {
	Ping() error
	CheckDb() error
	Db() (*sql.DB, error)
	Close()
	GetName() string
	IsCreated() bool
	SetCreated()
	Driver() string
}

// описатели БД
type Dbs interface {
	Znak() DbInfo
	Lite() DbInfo
	Config() DbInfo
	A3() DbInfo
	ZnakSvc() DbService
	LiteSvc() DbService
	ConfigSvc() DbService
	A3Svc() DbService
	SaveConfig(cfg IConfig) (err error)
}

type DbInfo interface {
	Name() string
	File() string
	Driver() string
	Uri() string
	Absent() bool // отсутствует нет пинга
	InfoDbService() DbService
	SaveConfig(cfg IConfig, key string) (err error)
	String() string
}
