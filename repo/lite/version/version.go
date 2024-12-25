package version

import (
	"fmt"
	"os"

	"firstwails/entity"
	"firstwails/repo/lite/version/create"
	v002 "firstwails/repo/lite/version/v002"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

type IRepo interface {
	Configuration() *entity.Configuration
	Logger() *zap.SugaredLogger
	Config() entity.IConfig
	LiteDbService() entity.DbService
}

const modError = "repo:newrequest:version"

// const setVersion string = `INSERT OR REPLACE INTO dboptions (name, value) VALUES ('version',?)`
const getVersionSql string = `select value from dboptions where name = 'version' limit 1;`

func CheckVersionDb(app IRepo) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
	}()
	dbSvc := app.LiteDbService()
	if dbSvc == nil {
		return fmt.Errorf("report:version не инициализирована служба DbService для БД report")
	}
	if dbSvc.IsCreated() {
		// файл БД только что создан и требуется инициализация полной структуры БД
		if err := create.New(app).Upgrade(); err != nil {
			// надо удалить плохо созданный БД
			name := dbSvc.GetName()
			os.Remove(name)
			return fmt.Errorf("repo:version новая БД %s удалена из за ошибки создания %v", name, err)
		}
	} else {
		// файл БД существует надо проверить версию и необходимость обновления структуры БД
		if versionInFile, err := getStringVersionDb(app); err != nil {
			return fmt.Errorf("repo:version %v", err)
		} else {
			// TODO сздесь проверка должна быть с v002.DBVersion ...
			if versionInFile == create.DBVersion {
				// версия текущая не надо ни чего делать
				return nil
			} else if versionInFile > create.DBVersion {
				// это будущая версия не верно ... надо выходить с ошибкой
				return fmt.Errorf("версия бд ошибочна обратитесь в службу поддержки")
			} else {
				switch versionInFile {
				case "0.0.1":
					// бэкап БД в каталог entity.BackUpPath
					// предыдущая версия, тут будут все версии что выходят
					if err := v002.New(app).Upgrade(); err != nil {
						return fmt.Errorf("v002 ошибка обновления %w", err)
					}
					return nil
				default:
					return fmt.Errorf("ошибка версии обратитесь в службу поддержки")
				}
			}
		}
	}
	return nil
}
func getStringVersionDb(r IRepo) (out string, err error) {
	var dbFunc *sqlx.DB
	if db, err := r.LiteDbService().Db(); err != nil {
		r.Logger().Errorf("repo:config error %s", err.Error())
		return "", fmt.Errorf("%w", err)
	} else {
		dbFunc = sqlx.NewDb(db, "sqlite3")
	}
	defer func() {
		dbFunc.Close()
		r.LiteDbService().Close()
	}()

	// по умолчанию берем последнюю версию
	// TODO по логике надо версию пакета обновления максимального...
	ver := create.DBVersion
	// проверяем есть ли вообще версия в БД
	if err := dbFunc.Get(&ver, getVersionSql); err != nil {
		return "", fmt.Errorf("db.Get() %w", err)
	}
	return ver, nil
}

// func getVersion(app entity.App) entity.VersionDB {
// 	return nil
// }
