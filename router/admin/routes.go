package router

import (
	"dropCms/handlers/pages"

	"github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(r fiber.Router) {
  r.Get("/", handlers.Index)
  r.Get("/login", handlers.Login)
}
