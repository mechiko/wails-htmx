package repo

import (
	"database/sql"
	"firstwails/migrations"
	"fmt"

	"github.com/pressly/goose/v3"
)

// при инициализации приложения этот метод вызывается однажды и прописывает объект доступа
// к базе данных, далее проверяет версию БД возможна ошибка и нужно выходить из приложения
func (r *repository) lite() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("repo:lite panic %v", r)
		}
	}()
	// return version.CheckVersionDb(r)
	if r.dbs.Lite().Absent() {
		return fmt.Errorf("database lite.db absent")
	}
	db, err := r.dbs.LiteSvc().Db()
	defer func() {
		db.Close()
		r.dbs.LiteSvc().Close()
	}()
	if err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	if err := makeMigrations(db); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	return nil
}

func makeMigrations(DB *sql.DB) error {
	goose.SetBaseFS(migrations.MigrationFiles)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}
	if err := goose.Up(DB, "."); err != nil {
		return err
	}
	return nil
}
