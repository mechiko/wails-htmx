package create

import (
	_ "embed"
	"fmt"

	"firstwails/domain"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// пакет сreate создает БД целиком содержит скрипты для полного создания
// версия БД зашита тут

const DBVersion = "0.0.1"

const modError = "version:create"

//go:embed whatsNew.txt
var whatsNew string

type versionDb struct {
	IRepo
	version string
	db      *sqlx.DB
}

var _ domain.VersionDB = (*versionDb)(nil)

type IRepo interface {
	Configuration() *domain.Configuration
	Logger() *zap.SugaredLogger
	LiteDbService() domain.IDbService
}

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

func (vv *versionDb) closeDB() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
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
		vv.closeDB()
		if r := recover(); r != nil {
			err = fmt.Errorf("report:upgrade panic %v", r)
		}
	}()

	if err := vv.create(); err != nil {
		return fmt.Errorf("report:upgrade %w", err)
	}
	// устанавливаем версию БД в конце Upgrade
	if _, err := vv.db.Exec(setVersion, DBVersion); err != nil {
		return fmt.Errorf("repo:create %w", err)
	}
	return err
}

func (vv *versionDb) Version() string {
	return vv.version
}

func (vv *versionDb) WhatsNew() string {
	return vv.version
}
