package main

import (
	migration "gofiber-restapi/db/migrations"
	seed "gofiber-restapi/db/seeders"
	"gofiber-restapi/internal/api"
	"gofiber-restapi/internal/config"
	"gofiber-restapi/internal/connection"
	"gofiber-restapi/internal/repository"
	"gofiber-restapi/internal/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	if err := migration.Run(dbConnection); err != nil {
		log.Fatal(err)
	}
	if os.Getenv("RUN_SEEDER") == "true" {
		if err := seed.Run(dbConnection); err != nil {
			log.Fatal("Seeder failed:", err)
		}
	}

	customerRepository := repository.NewCustomer(dbConnection)
	customerService := services.NewCustomer(customerRepository)

	app := fiber.New()

	api.NewCustomer(app, customerService)
	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
