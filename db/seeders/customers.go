package seed

import (
	"database/sql"
	"log"
)

func SeedCustomers(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM customers").Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("SeedCustomers: data already exists, skipping")
		return nil
	}

	_, err = db.Exec(`
		INSERT INTO customers (id, code, name, created_at, updated_at, deleted_at) VALUES
		('11111111-1111-1111-1111-111111111111', 'CUST001', 'Customer One', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
		('22222222-2222-2222-2222-222222222222', 'CUST002', 'Customer Two', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL)
	`)
	return err
}
