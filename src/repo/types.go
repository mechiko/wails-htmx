package repo

import (
	"fmt"
	"strings"

	"firstwails/domain"
	"firstwails/repo/a3/a3mssql"
	"firstwails/repo/a3/a3sqlite"
	"firstwails/repo/config"
	"firstwails/repo/dbs"
	"firstwails/repo/lite"
	"firstwails/repo/znak/znakmssql"
	"firstwails/repo/znak/znaksqlite"
	"firstwails/utility"
)

const modError = "pkg:repo"

// в каждом пакете этого модуля (dbconfig, dvreport ...)
// создается структура с интерфейсом domain для выполнения операция по БД
// чтобы управлять блокировками для операция в БД создаем в каждом модуле
// метод CloseDB() для закрытия ресурса и его надо выполнять defer в каждой функции вызова метода
// всем методы должны быть разовыми при вызове создает объект интерфейс и открывается БД

type repository struct {
	domain.IApp
	dbs domain.Dbs
}

// для создания репозитория требуется объект с интерфейсом IApp
// type IApp interface {
// 	Configuration() *domain.Configuration
// 	Logger() *zap.SugaredLogger
// 	Config() domain.IConfig
// 	Pwd() string
// 	FsrarID() string
// 	StartDateString() string
// 	EndDateString() string
// 	SetFsrarID(string)
// }

var _ domain.Repo = &repository{}

func New(a domain.IApp) (rr *repository, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	rp := &repository{
		IApp: a,
	}
	// создаем объект описателей БД
	// имя БД конфигурации А3 и с признаком сканирования других БД
	rp.dbs = dbs.New(a.Configuration(), "config.db", true)
	if strings.ToLower(rp.dbs.Lite().Driver()) != "sqlite" {
		panic("программа поддерживает только SQLITE базу данных")
	}

	if rp.dbs.A3().Absent() {
		// return nil, fmt.Errorf("%s отсутствует БД А3 драйвер %s", modError, rp.dbs.A3().Driver())
		rp.Logger().Infof("%s отсутствует БД А3 [драйвер %s]", modError, rp.dbs.A3().Driver())
	} else {
		// если драйвер mssql то конфиг фсрар ид устанавливаем по имени БД
		if strings.ToLower(rp.dbs.A3().Driver()) == "mssql" {
			if a.Configuration().Application.Fsrarid != rp.dbs.A3().Name() {
				a.Config().Set("application.fsrarid", rp.dbs.A3().Name(), true)
			}
		}
		if strings.ToLower(rp.dbs.A3().Driver()) == "sqlite" {
			dbname := utility.FileNameWithoutExtension(rp.dbs.A3().File())
			if a.Configuration().Application.Fsrarid != dbname {
				a.Config().Set("application.fsrarid", dbname, true)
			}
		}
	}
	if rp.dbs.Config().Absent() {
		// return nil, fmt.Errorf("%s отсутствует БД config драйвер %s", modError, rp.dbs.Config().Driver())
		rp.Logger().Infof("%s отсутствует БД config драйвер %s", modError, rp.dbs.A3().Driver())
	}
	if rp.dbs.Lite().Absent() {
		// return nil, fmt.Errorf("%s отсутствует БД alcogo4 драйвер %s", modError, rp.dbs.Lite().Driver())
		rp.Logger().Infof("%s отсутствует БД lite драйвер %s", modError, rp.dbs.A3().Driver())
	}
	if rp.dbs.Znak().Absent() {
		// return nil, fmt.Errorf("%s отсутствует БД 4z драйвер %s", modError, rp.dbs.Znak().Driver())
		rp.Logger().Infof("%s отсутствует БД 4z драйвер %s", modError, rp.dbs.A3().Driver())
	}
	if err := rp.dbs.SaveConfig(a.Config()); err != nil {
		return nil, fmt.Errorf("%s %w", modError, err)
	}
	// инициализируем или проверяем БД Lite
	if err := rp.lite(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if rp.Configuration().Application.Fsrarid == "" {
		if rp.dbs.A3().Name() != "" {
			fsrarId := rp.dbs.A3().Name()
			if rp.dbs.A3().Driver() == "sqlite" {
				file := rp.dbs.A3().File()
				fsrarId = file[0 : len(file)-3]
				// fsrarId, _ = strings.CutSuffix(rp.dbs.A3().File(), ".db")
			}
			rp.SetFsrarID(fsrarId)
		}
	}
	return rp, nil
}

// метод возвращает из dbs.Znak описатель доступа к БД
// используется в конструкторе модуля repo/sqlite/znak типа domain.DbZnak
// в конструкторе вызывается открытие БД и по завершению использования надо по
// defer вызывать метод closeDB()
func (r *repository) ZnakDbService() domain.IDbService {
	return r.dbs.ZnakSvc()
}

// метод возвращает из dbs.Znak описатель доступа к БД
// используется в конструкторе модуля repo/sqlite/lite типа domain.DbLite
func (r *repository) LiteDbService() domain.IDbService {
	return r.dbs.LiteSvc()
}

func (r *repository) A3DbService() domain.IDbService {
	return r.dbs.A3Svc()
}

func (r *repository) DbLite() domain.DbLite {
	if r.dbs.Lite().Absent() {
		panic("отсутствует БД lite")
	}
	return lite.New(r)
}

func (r *repository) DbZnak() domain.DbZnak {
	if r.dbs.Znak().Absent() {
		panic("отсутствует БД 4z")
	}
	driver := r.dbs.Znak().Driver()
	if driver == "sqlite" {
		return znaksqlite.New(r)
	}
	if driver == "mssql" {
		return znakmssql.New(r)
	}
	return nil
}

func (r *repository) DbA3() domain.DbAlcohelp3 {
	if r.dbs.A3().Absent() {
		panic("отсутствует БД A3")
	}
	driver := r.dbs.A3().Driver()
	if driver == "sqlite" {
		return a3sqlite.New(r)
	}
	if driver == "mssql" {
		return a3mssql.New(r)
	}
	return nil
}

func (r *repository) Dbs() domain.Dbs {
	return r.dbs
}

func (r *repository) ConfigDbService() domain.IDbService {
	return r.dbs.ConfigSvc()
}

func (r *repository) DbConfig() domain.DbConfig {
	if r.dbs.Config().Absent() {
		panic("отсутствует БД config.db")
	}
	driver := r.dbs.Config().Driver()
	if driver != "sqlite" {
		panic("ошибочный драйвер БД config.db")
	}
	return config.New(r)
}
