package main

import (
	"test-api/app/services"
	"test-api/migrations"
	"test-api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	services.InitDatabase()
	migrations.RunMigration()
	migrations.DataSeeds()

	router.Route(app)
	app.Listen(":3000")
}
