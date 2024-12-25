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

type dbService struct {
	ctx           context.Context
	sem           Semaphore
	connectionUri string
	driver        string
	dbName        string
	mode          UriType
	version       CheckMode
	created       bool
}

type Semaphore interface {
	Acquire()
	Release()
}
type semaphore struct {
	semC chan struct{}
}

func New(maxConcurrency int) Semaphore {
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

var _ domain.DbService = &dbService{}

// dbs, err := sqlite3.NewDbService(cfg.Database.DbName, sqlite3.RwModeWithCreate, sqlite3.Version)
// sqlite3.RwModeWithCreate константа в enum_types.go (sqlite3.RwModeWithCreate sqlite3.RoMode sqlite3.RwMode)
// versioning CheckMode сюда передаем sqlite3.Version если версионность поддерживается,
//
//	и sqlite3.NoMatter если не имеет значения
func NewDbService(dbname string, mode UriType, versioning CheckMode) (domain.DbService, error) {
	s := &dbService{
		dbName:        dbname,
		driver:        "sqlite3",
		mode:          mode,
		connectionUri: "", // прописываем если только создаем из URI
		version:       versioning,
		created:       false,
		ctx:           context.TODO(),
		sem:           New(maxWorkers),
	}
	if err := s.CheckDb(); err != nil {
		return s, fmt.Errorf("%s CheckDb %w", modError, err)
	}
	return s, nil
}

func NewFromUri(dbname string, uri string, versioning CheckMode) (domain.DbService, error) {
	s := &dbService{
		dbName:        dbname, // full path or relative
		driver:        "sqlite3",
		mode:          RoMode, // по умолчанию RO режим
		connectionUri: uri,    // используется для открытия БД если "" то будет нужен mode и будет фиг знает чего...
		version:       versioning,
		created:       false,
		ctx:           context.TODO(),
		sem:           New(maxWorkers),
	}
	if err := s.CheckDb(); err != nil {
		return s, fmt.Errorf("%s CheckDb %w", modError, err)
	}
	return s, nil
}

func (s *dbService) GetName() string {
	return s.dbName
}

func (s *dbService) Ping() error {
	if err := s.CheckDb(); err != nil {
		return fmt.Errorf("%s ping %w", modError, err)
	}
	db, err := sqlx.Open(s.driver, s.dbName+s.connectionUri)
	if err != nil {
		return fmt.Errorf("%s файл %s не существует %w", modError, s.dbName, err)
	}
	defer func() {
		db.Close()
	}()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("%s ping %w", modError, err)
	}

	return nil
}

func (s *dbService) CheckDb() error {
	s.created = false
	_, err := os.Stat(s.dbName)
	if os.IsNotExist(err) {
		if s.mode == RwModeWithCreate {
			if err := s.create(); err != nil {
				return fmt.Errorf("%s create %w", modError, err)
			}
			// создали без ошибки
			s.created = true
			return nil
		} else {
			// открытие без создания и файла нет
			return fmt.Errorf("%s file not exists %w", modError, err)
		}
	} else if err != nil {
		// какая то другая ошибка кроме наличия файла
		return fmt.Errorf("%s файл %s ошибка %w", modError, s.dbName, err)
	}
	return nil
}

func (s *dbService) Db() (*sql.DB, error) {
	s.sem.Acquire()
	err := s.CheckDb()
	if err != nil {
		return nil, fmt.Errorf("%s CheckDb %w", modError, err)
	}
	dbUri := fmt.Sprintf("%s%s", s.dbName, s.mode.String())
	if s.connectionUri != "" {
		dbUri = fmt.Sprintf("%s%s", s.dbName, s.connectionUri)
	}
	dbConn, err := sql.Open(s.driver, dbUri)
	if err != nil {
		return nil, fmt.Errorf("%s %w", modError, err)
	}
	dbConn.SetMaxOpenConns(0)
	return dbConn, nil
}

func (s *dbService) Close() {
	// fmt.Printf("Close dbService DB file:%s\n\r", s.DbName)
	s.sem.Release()
}

// func (s *dbService)	GetRequestRules() (*domain.RequestRuleList, error) {
// }
// в момент открыти создается файл БД
func (s *dbService) create() error {
	db, err := sqlx.Open(s.driver, s.dbName+s.connectionUri)
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

func (s *dbService) Driver() string {
	return s.driver
}
