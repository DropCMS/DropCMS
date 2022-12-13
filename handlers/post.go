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
			"err": parseErr.Error(),
		})
	}
	tokens := c.Locals("users").(*jwt.Token)
	claims := tokens.Claims.(jwt.MapClaims)
	post.UserId = claims["user_id"].(uint)
	posId, err := post.Create()
	postId := posId.(int64)
	if err != nil {
		return c.JSON(fiber.Map{
			"err": err.Error(),
		})
	}
	var tagId int64
	for _, v := range post.Tags {
		tag := new(models.Tag)
		tag.Name = v
		id, erru := tag.Create()
		tagId = id.(int64)
		if erru != nil {
			return c.JSON(fiber.Map{
				"err": erru.Error(),
			})
		}
		erur := models.MakeAssociationWithTag(uint(postId), uint(tagId))
		if erur != nil {
			return c.JSON(fiber.Map{
				"err": erur.Error(),
			})
		}
	}
	return c.JSON(fiber.Map{
		"message": "Post created.",
	})
}
