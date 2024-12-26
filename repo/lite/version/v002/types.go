package v002

import (
	_ "embed"
	"fmt"

	"firstwails/domain"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// пакет сreate создает БД целиком содержит скрипты для полного создания
// версия БД зашита тут

const DBVersion = "0.0.2"

const modError = "version:v002"

//go:embed whatsNew.txt
var whatsNew string

type versionDb struct {
	IRepo
	db      *sqlx.DB
	version string
}
type IRepo interface {
	Configuration() *domain.Configuration
	Logger() *zap.SugaredLogger
	// BackupFile(string) error
	LiteDbService() domain.IDbService
}

var _ domain.VersionDB = (*versionDb)(nil)

// создаем структуру с номером версии и методами для создания
func New(app IRepo) *versionDb {
	rc := &versionDb{
		IRepo:   app,
		version: DBVersion,
	}
	if db, err := app.LiteDbService().Db(); err != nil {
		app.Logger().Errorf("%s %s", modError, err.Error())
	} else {
		rc.db = sqlx.NewDb(db, "sqlite3")
	}
	return rc
}

func (vv *versionDb) сloseDB() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	// это может вернуть ошибку даже если второй вызов будет без ошибки
	err = vv.db.Close()
	vv.LiteDbService().Close()
	return err
}

// в пакете create это будет созданием БД новой
// со структурой на текущий момент
func (vv *versionDb) Upgrade() (err error) {
	defer func() {
		err = vv.сloseDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("report:upgrade panic %v", r)
		}
	}()

	// перед обновлением делаем бэкап при невозможности не продолжаем выходим по ошибке
	// name := vv.DriverReportDb().GetName()
	// if err := vv.BackupFile(name); err != nil {
	// 	return fmt.Errorf("repo:upgrade backup %w", err)
	// }
	if err := vv.create(); err != nil {
		return fmt.Errorf("repo:upgrade %w", err)
	}
	// патчим версию если не было ошибок
	if _, err = vv.db.Exec(setVersion, DBVersion); err != nil {
		return fmt.Errorf("repo:upgrade %w", err)
	}
	return err
}

func (vv *versionDb) Version() string {
	return vv.version
}

func (vv *versionDb) WhatsNew() string {
	return vv.version
}
