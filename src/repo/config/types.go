package config

import (
	"database/sql"
	_ "embed"
	"fmt"

	"firstwails/domain"

	"go.uber.org/zap"
)

type dbConfig struct {
	IRepo
	db *sql.DB
}

type IRepo interface {
	Configuration() *domain.Configuration
	Logger() *zap.SugaredLogger
	Config() domain.IConfig
	FsrarID() string
	StartDateString() string
	EndDateString() string

	ConfigDbService() domain.IDbService
}

const modError = "repo:a3sqlite:dbalcohelp3"

func New(a IRepo) domain.DbConfig {
	var err error
	rc := &dbConfig{
		IRepo: a,
	}
	// if rc.db, err = a.DriverAlcoHelpDb().Db(); err != nil {
	// 	a.GetLogger().Errorf("repo:alcohelp error %s", err.Error())
	// }
	if rc.db, err = a.ConfigDbService().Db(); err != nil {
		// a.GetLogger().Errorf("%s error %s", modError, err.Error())
		panic(fmt.Errorf("%s error %s", modError, err.Error()))
	}

	return rc
}

func (c *dbConfig) сloseDB() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	// для sqlite вызываем закрыти чтобы освободить семафор
	err = c.db.Close()
	c.ConfigDbService().Close()
	return err
}
