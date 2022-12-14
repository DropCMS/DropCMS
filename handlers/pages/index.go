package handlers

import (
	"dropCms/utils"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
  token := c.Cookies("access_token")
  tk := utils.Parse(token, c)
  if tk == nil {
    return c.Redirect("/login")
  }
  return c.Render("admin/main", fiber.Map{},"layout/main")
}
