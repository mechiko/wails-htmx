package dbsrc

import (
	"log"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	_ "modernc.org/sqlite"
)

type dbsrc struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) *dbsrc {
	db, err := sqlx.Open("sqlite", `M:\!AlcoSoft\A3 UTSZ\030000527832.db`)
	// db, err := sqlx.Open("sqlite", `C:\!DB\UTSZ\030000527832.db`)
	if err != nil {
		log.Fatal(err)
	}
	return &dbsrc{
		db: db,
	}
}

func (a *dbsrc) Close() error {
	return a.db.Close()
}
