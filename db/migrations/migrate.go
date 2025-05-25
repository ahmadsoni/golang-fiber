package migration

import (
	"database/sql"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Run(db *sql.DB) error {
	log.Println("[Migration] Starting migration...")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.New("[Migration] Failed to create migration driver: " + err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return errors.New("[Migration] Migration setup failed: " + err.Error())
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("[Migration] No new migration to apply.")
		} else {
			return errors.New("[Migration] Migration failed: " + err.Error())
		}
	} else {
		log.Println("[Migration] Migration applied successfully.")
	}

	return nil
}
