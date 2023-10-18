package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/mdjarv/templcards/pkg/database"
	"github.com/mdjarv/templcards/pkg/routes"
)

func main() {
	db, err := database.New()
	if err != nil {
		panic("failed to connect database")
	}
	println("Connected to database")

	err = db.Migrate()
	if err != nil {
		panic("failed to migrate database")
	}
	println("Migrated database")

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/metrics", monitor.New())

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString(("OK"))
	})

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: http.Dir("./assets"),
	}))

	routes.Setup(app, db)

	log.Fatal(app.Listen(":8080"))
}
