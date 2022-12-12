package router

import (
	"dropCms/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(f fiber.Router) {
	f.Post("/register", handlers.RegisterUser)
	f.Post("/login", handlers.LoginUser)
}
