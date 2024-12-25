package znaksqlite

import (
	_ "embed"
	"fmt"

	"firstwails/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const modError = "repo:znak"

type dbZnak struct {
	IRepo
	db *sqlx.DB
}

type IRepo interface {
	Configuration() *entity.Configuration
	Logger() *zap.SugaredLogger
	Config() entity.IConfig

	ZnakDbService() entity.DbService
}

// вызывается каждый раз когда требуется доступ к БД конфиг
func New(r IRepo) entity.DbZnak {
	rc := &dbZnak{
		IRepo: r,
	}
	if db, err := r.ZnakDbService().Db(); err != nil {
		panic(fmt.Errorf("%s New panic %s", modError, err.Error()))
	} else {
		rc.db = sqlx.NewDb(db, "sqlite3")
	}
	return rc
}

// вызывается defer внутри модуля
func (c *dbZnak) сloseDB() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s сloseDB panic %v", modError, r)
		}
	}()
	// это может вернуть ошибку даже если второй вызов будет без ошибки
	// закрываем БД
	err = c.db.Close()
	// освобождает семафор для БД
	c.ZnakDbService().Close()
	return err
}
