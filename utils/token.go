package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var SuccessHandler fiber.Handler
var ErrHandler fiber.Handler

func Authorize(c *fiber.Ctx) error {
	hh := c.Cookies("access_token")
	if hh == "" {
		return c.Redirect("/admin/login")
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
	return c.Redirect("/admin/login")
}

func Parse(token string, c *fiber.Ctx) (tt *jwt.Token) {
  toke, errorr := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
    _,ok := t.Method.(*jwt.SigningMethodHMAC)

    if !ok {
      return nil, errors.New("wrong parsing method")
    }
    return []byte("mypasswordaaltufaltukhaopio"), nil
  })
  if errorr != nil {
    c.JSON(fiber.Map{
      "err" : errorr.Error(),
    })
  }
  if toke.Valid {
    return toke
  }
  return nil
} 
