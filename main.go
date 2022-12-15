package main

import (
	"dropCms/db"
	router "dropCms/router/api"
  routes "dropCms/router/admin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	db.ConnectDb()

	api := app.Group("/api")
	router.Routes(api)
  admin := app.Group("/admin")
  routes.SetupAdminRoutes(admin)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Srt")
	})
	app.Listen(":8080")
}
