package lite

import (
	_ "embed"
	"fmt"

	"firstwails/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const modError = "repo:lite"

type dbLite struct {
	IRepo
	db *sqlx.DB
}

type IRepo interface {
	Configuration() *entity.Configuration
	Logger() *zap.SugaredLogger
	Config() entity.IConfig

	LiteDbService() entity.DbService
}

// вызывается каждый раз когда требуется доступ к БД конфиг
func New(r IRepo) entity.DbLite {
	rc := &dbLite{
		IRepo: r,
	}
	if db, err := r.LiteDbService().Db(); err != nil {
		r.Logger().Errorf("repo:config error %s", err.Error())
		panic(fmt.Errorf("%s error %s", modError, err.Error()))
	} else {
		rc.db = sqlx.NewDb(db, "sqlite3")
	}
	return rc
}

// вызывается defer внутри модуля
func (c *dbLite) сloseDB() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	// это может вернуть ошибку даже если второй вызов будет без ошибки
	// закрываем БД
	err = c.db.Close()
	// освобождает семафор для БД
	c.LiteDbService().Close()
	return err
}
