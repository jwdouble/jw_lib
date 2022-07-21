package migrate

import (
	"database/sql"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"jw.lib/logx"
)

func NewWithFs(fs fs.FS, path string, db *sql.DB) {
	si, err := iofs.New(fs, path)
	if err != nil {
		logx.Errorf(err, "migrate.NewWithFs")
		return
	}

	di, err := postgres.WithInstance(db, &postgres.Config{MigrationsTable: "_migration"})
	if err != nil {
		logx.Errorf(err, "postgres.WithInstance")
		return
	}

	m, err := migrate.NewWithInstance("iofs", si, "postgres", di)
	if err != nil {
		logx.Errorf(err, "migrate.NewWithInstance")
		return
	}

	m.Log = logger{}

	err = m.Up()
	if err != nil {
		logx.Errorf(err, "migrate.Up")
	}
}

type logger struct {
}

func (l logger) Printf(format string, v ...interface{}) {
	logx.Debugf(format, v...)
}

func (l logger) Verbose() bool {
	return true
}
