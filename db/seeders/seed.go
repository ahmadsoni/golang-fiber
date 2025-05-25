package seed

import (
	"database/sql"
)

func Run(db *sql.DB) error {
	if err := SeedCustomers(db); err != nil {
		return err
	}
	return nil
}
