package migrations

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

const dialect = "postgres"

//go:embed *.sql
var embedMigrations embed.FS

func MaybeApplyMigrations(dbHandle *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	if err := goose.Up(dbHandle, "."); err != nil {
		return err
	}

	return nil
}
