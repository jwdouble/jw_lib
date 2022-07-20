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
		logx.Error("iofs.New", err)
		return
	}

	di, err := postgres.WithInstance(db, &postgres.Config{MigrationsTable: "_migration"})
	if err != nil {
		logx.Error("stub.WithInstance", err)
		return
	}

	m, err := migrate.NewWithInstance("iofs", si, "postgres", di)
	if err != nil {
		logx.Error("migrate.NewWithDatabaseInstance", err)
		return
	}

	err = m.Up()
	if err != nil {
		logx.Error("migrate.Up", err)
	}
}
