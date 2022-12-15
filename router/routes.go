package router

import (
	handlers "dropCms/handlers/pages"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
  app.Get("/login",handlers.Login)
}
