package utils

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Authorize(c *fiber.Ctx) error {
	hh := c.Cookies("access_token")
	if hh == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	token, err := jwt.Parse(hh, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("wrong parsing method.")
		}
		return []byte("mypasswordaaltufaltukhaopio"), nil
	})
	if err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}
	if token.Valid {
		c.Locals("users", token)
		return c.Next()
	}
	return c.SendStatus(fiber.StatusForbidden)
}

func AuthorizeFromHeader(c *fiber.Ctx) error {
	hh := c.Get("Authorization")
	ll := len("Bearer")
	if len(hh) > ll+1 && strings.EqualFold(hh[:ll], "Bearer") {
		tkn := strings.TrimSpace(hh[ll:])
		token, jtErr := jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("wrong parsing method.")
			}
			return []byte("mypasswordaaltufaltukhaopio"), nil

		})
		if jtErr != nil {
			return c.JSON(fiber.Map{"err": jtErr.Error()})
		}
		if token.Valid {
			c.Locals("users", token)
			return c.Next()
		}
		return c.SendStatus(fiber.StatusForbidden)

	}
	return c.SendStatus(fiber.StatusUnauthorized)
}
