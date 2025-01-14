package domain

import "database/sql"

type IDbService interface {
	Ping() error
	CheckDb() error
	Db() (*sql.DB, error)
	Close()
	DbInfo() DbInfo
	IsCreated() bool
	SetCreated()
}

// описатели БД
type Dbs interface {
	Znak() DbInfo
	Lite() DbInfo
	Config() DbInfo
	A3() DbInfo
	ZnakSvc() IDbService
	LiteSvc() IDbService
	ConfigSvc() IDbService
	A3Svc() IDbService
	SaveConfig(cfg IConfig) (err error)
}

type DbInfo interface {
	Name() string
	File() string
	Driver() string
	Absent() bool // отсутствует нет пинга
	InfoDbService() IDbService
	SaveConfig(cfg IConfig, key string) (err error)
	String() string
	Mode() UriMode
	CheckMode() CheckMode
	URI() string
}

type UriMode int

const (
	RwModeWithCreate UriMode = iota
	RoMode
	RwMode
)

func (s UriMode) String() string {
	switch s {
	case RwModeWithCreate:
		return "mode=rwc&_journal_mode=WAL"
	case RoMode:
		return "mode=ro"
	case RwMode:
		return "mode=rw&_journal_mode=WAL"
	}
	return "mode=rw&_journal_mode=WAL"
}

type CheckMode int

const (
	NoMatter CheckMode = iota
	Version
	MinVersion
)

func (s CheckMode) String() string {
	switch s {
	case Version:
		return "версия проверятся c миграцией"
	case MinVersion:
		return "версия проверятся на минимальное значение"
	case NoMatter:
		return "ни каких проверок"
	}
	return "не проверятся"
}
