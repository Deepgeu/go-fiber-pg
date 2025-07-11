package main

import (
	"go-fiber-pg/db"
	"go-fiber-pg/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(" ** Server is running ** ")
	})

	app.Get("/seed", handlers.SeedRecords)
	app.Get("/records", handlers.GetAllRecords)
	app.Post("/record", handlers.CreateRecord)
	app.Get("/record/:key", handlers.GetRecord)
	app.Put("/record/:key", handlers.UpdateRecord)
	app.Delete("/record/:key", handlers.DeleteRecord)

	app.Listen(":3000")
}
