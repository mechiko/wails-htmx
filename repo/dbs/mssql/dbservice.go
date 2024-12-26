// Package postgres implements postgres connection.
package mssql

import (
	"context"
	"database/sql"
	"fmt"

	"firstwails/domain"

	_ "github.com/denisenkom/go-mssqldb"
)

type dbService struct {
	ctx    context.Context
	dbInfo domain.DbInfo
	dBConn *sql.DB
}

type IDbService interface {
	Ping() error
	CheckDb() error
	Db() (*sql.DB, error)
	Close()
	DbInfo() domain.DbInfo
	IsCreated() bool
	SetCreated()
}

const modError = "dbs:dbservice:mssql"
const driver = "mssql"

var _ domain.IDbService = &dbService{}
var _ IDbService = &dbService{}

func New(info domain.DbInfo) (ss domain.IDbService, err error) {
	s := &dbService{
		dbInfo: info,
		ctx:    context.TODO(),
	}
	if err := s.CheckDb(); err != nil {
		return s, fmt.Errorf("%s %w", modError, err)
	}
	if err := s.Ping(); err != nil {
		return s, fmt.Errorf("%s %w", modError, err)
	}
	return s, nil
}

func (s *dbService) Ping() error {
	if s.dBConn == nil {
		return fmt.Errorf("%s %s", modError, "dbconn is nil")
	}
	if err := s.dBConn.Ping(); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}

	return nil
}

func (s *dbService) CheckDb() (err error) {
	s.dBConn, err = sql.Open("mssql", s.dbInfo.URI())
	return err
}

func (s *dbService) Db() (*sql.DB, error) {
	if err := s.Ping(); err != nil {
		return s.dBConn, fmt.Errorf("%s %w", modError, err)
	}
	return s.dBConn, nil
}

func (s *dbService) Close() {
}

func (s *dbService) create() error {
	return nil
}

func (s *dbService) IsCreated() bool {
	return false
}

func (s *dbService) SetCreated() {
}

func (s *dbService) DbInfo() domain.DbInfo {
	return s.dbInfo
}
