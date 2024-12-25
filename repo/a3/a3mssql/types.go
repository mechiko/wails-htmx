package a3mssql

import (
	"database/sql"
	_ "embed"
	"fmt"

	"firstwails/entity"

	"go.uber.org/zap"
)

type dbAlcohelp3 struct {
	IRepo
	db *sql.DB
}

type IRepo interface {
	Configuration() *entity.Configuration
	Logger() *zap.SugaredLogger
	Config() entity.IConfig
	FsrarID() string
	StartDateString() string
	EndDateString() string

	A3DbService() entity.DbService
}

const modError = "repo:a3mssql:dbalcohelp3"

func New(a IRepo) entity.DbAlcohelp3 {
	var err error
	rc := &dbAlcohelp3{
		IRepo: a,
	}
	if rc.db, err = a.A3DbService().Db(); err != nil {
		panic(fmt.Errorf("%s error %s", modError, err.Error()))
	}

	return rc
}

func (c *dbAlcohelp3) сloseDB() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	// для mssql закрывать не надо думаю
	// err = c.db.Close()
	return err
}
