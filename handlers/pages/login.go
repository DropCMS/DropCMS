package handlers

import (
	"dropCms/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
  token := c.Cookies("access_token")
  if token != "" { 
  tt := utils.Parse(token, c)
  if tt != nil {
    return c.Redirect("/")
  }
  }
  return c.Render("admin/login", fiber.Map{},
  "layout/main")
}
