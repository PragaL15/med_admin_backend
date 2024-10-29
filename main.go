package main

import (
	"github.com/PragaL15/med_admin_backend/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitializeDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
