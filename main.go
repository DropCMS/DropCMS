package main

import (
	"dropCms/db"
	"dropCms/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db.ConnectDb()
	api := app.Group("/api")
	router.Routes(api)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Srt")
	})
	app.Listen(":8080")
}
