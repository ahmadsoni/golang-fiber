package seed

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SeedUsers(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("SeedUsers: data already exists, skipping")
		return nil
	}

	// Hash passwords
	userPassword, err := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	adminPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	INSERT INTO users (id, name, email, password, created_at, updated_at, deleted_at) VALUES
	('33333333-3333-3333-3333-333333333333', 'User Biasa', 'user@server-sun.my.id', $1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
	('44444444-4444-4444-4444-444444444444', 'Admin Utama', 'admin@server-sun.my.id', $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL)
`, string(userPassword), string(adminPassword))

	if err != nil {
		return err
	}

	log.Println("SeedUsers: users seeded successfully")
	return nil
}
