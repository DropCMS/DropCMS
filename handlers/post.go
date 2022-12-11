package handlers

import (
	"dropCms/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreatePost(c *fiber.Ctx) error {
  post := new(models.Post)
  parseErr := c.BodyParser(post)
  if parseErr != nil {
    return c.JSON(fiber.Map{
      "err" : parseErr.Error(),
    })
  }
  tokens := c.Locals("users").(*jwt.Token)
  claims := tokens.Claims.(jwt.MapClaims)
  post.UserId = claims["user_id"].(uint)
  return c.JSON(fiber.Map{
    "message" : "Post created.",
  })
}
