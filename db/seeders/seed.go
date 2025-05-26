package seed

import (
	"database/sql"
	"log"
)

func Run(db *sql.DB) error {
	log.Println("Seeder: start seeding...")

	if err := SeedCustomers(db); err != nil {
		log.Println("Seeder: failed to seed customers:", err)
		return err
	}

	if err := SeedUsers(db); err != nil {
		log.Println("Seeder: failed to seed users:", err)
		return err
	}

	log.Println("Seeder: seeding completed successfully")
	return nil
}
