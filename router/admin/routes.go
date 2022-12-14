package router

import "github.com/gofiber/fiber/v2"

func SetupAdminRoutes(r fiber.Router) {
  r.Get("/");
}
