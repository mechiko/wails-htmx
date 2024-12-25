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
	ctx           context.Context
	connectionUri string
	driver        string
	dbName        string
	// user          string
	// host          string
	// pass          string
	dBConn *sql.DB
}

const modError = "dbs:mssql"

var _ domain.DbService = &dbService{}

func New(dbname string, uri string) (ss domain.DbService, err error) {
	s := &dbService{
		dbName:        dbname,
		driver:        "mssql",
		connectionUri: uri + fmt.Sprintf(";database=%s", dbname),
		ctx:           context.TODO(),
	}
	if err := s.CheckDb(); err != nil {
		return s, fmt.Errorf("%s %w", modError, err)
	}
	if err := s.Ping(); err != nil {
		return s, fmt.Errorf("%s %w", modError, err)
	}
	return s, nil
}

func NewFromUri(dbname string, uri string) (ss domain.DbService, err error) {
	s := &dbService{
		dbName:        dbname,
		driver:        "mssql",
		connectionUri: uri,
		ctx:           context.TODO(),
	}
	if err := s.CheckDb(); err != nil {
		return s, fmt.Errorf("%s %w", modError, err)
	}
	if err := s.Ping(); err != nil {
		return s, fmt.Errorf("%s %w", modError, err)
	}
	return s, nil
}

func (s *dbService) GetName() string {
	return s.dbName
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
	s.dBConn, err = sql.Open("mssql", s.connectionUri)
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

func (s *dbService) Driver() string {
	return s.driver
}
