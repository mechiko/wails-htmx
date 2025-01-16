package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"firstwails/domain"

	"github.com/jmoiron/sqlx"

	// _ "modernc.org/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

const modError = "dbs:dbservice:sqlite3"

const driver = "sqlite3"

type IDbService interface {
	Ping() error
	CheckDb() error
	Db() (*sql.DB, error)
	Close()
	DbInfo() domain.DbInfo
	IsCreated() bool
	SetCreated()
}

type dbService struct {
	ctx     context.Context
	sem     Semaphore
	dbInfo  domain.DbInfo
	created bool
}

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semC chan struct{}
}

func NewSemaphore(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}

func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semC
}

// глубина канала сколько одновременно может работать с БД всегда с sqlite один!
const maxWorkers = 1

var _ IDbService = &dbService{}
var _ domain.IDbService = &dbService{}

func New(info domain.DbInfo) (domain.IDbService, error) {
	if info == nil {
		panic(fmt.Sprintf("%s dbinfo is nil", modError))
	}
	s := &dbService{
		dbInfo:  info,
		created: false,
		ctx:     context.TODO(),
		sem:     NewSemaphore(maxWorkers),
	}
	if err := s.CheckDb(); err != nil {
		return s, fmt.Errorf("%s CheckDb %w", modError, err)
	}
	return s, nil
}

func (s *dbService) Ping() error {
	if err := s.CheckDb(); err != nil {
		return fmt.Errorf("%s ping %w", modError, err)
	}
	db, err := sqlx.Open(driver, s.dbInfo.URI())
	if err != nil {
		return fmt.Errorf("%s файл %s ошибка открытия %w", modError, s.dbInfo.File(), err)
	}
	defer func() {
		db.Close()
	}()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("%s ping %w", modError, err)
	}

	return nil
}

// проверка наличия файла физически
// если файла нет и режим с созданием то открываем через драйвер
func (s *dbService) CheckDb() error {
	s.created = false
	_, err := os.Stat(s.dbInfo.File())
	if os.IsNotExist(err) {
		if s.dbInfo.Mode() == domain.RwModeWithCreate {
			if err := s.create(); err != nil {
				return fmt.Errorf("%s create %w", modError, err)
			}
			// создали без ошибки
			s.created = true
			return nil
		} else {
			// открытие без создания и файла нет
			return fmt.Errorf("%s file %s not exists %w", modError, s.dbInfo.File(), err)
		}
	} else if err != nil {
		// какая то другая ошибка кроме наличия файла
		return fmt.Errorf("%s файл %s ошибка %w", modError, s.dbInfo.File(), err)
	}
	return nil
}

func (s *dbService) Db() (*sql.DB, error) {
	s.sem.Acquire()
	dbConn, err := sql.Open(driver, s.dbInfo.URI())
	if err != nil {
		return nil, fmt.Errorf("%s %w", modError, err)
	}
	dbConn.SetMaxOpenConns(0)
	return dbConn, nil
}

func (s *dbService) Close() {
	s.sem.Release()
}

// в момент открыти создается файл БД
func (s *dbService) create() error {
	db, err := sql.Open(driver, s.dbInfo.URI())
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	return db.Close()
}

func (s *dbService) IsCreated() bool {
	return s.created
}

func (s *dbService) SetCreated() {
	s.created = true
}

func (s *dbService) DbInfo() domain.DbInfo {
	return s.dbInfo
}
