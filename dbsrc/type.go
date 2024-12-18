package dbsrc

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type dbsrc struct {
	db     *sqlx.DB
	logger *log.Logger
}

func New(logger *log.Logger) *dbsrc {
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
