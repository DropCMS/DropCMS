package handlers

import (
	"database/sql"
	"dropCms/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(&user)
	if err != nil {
		return c.JSON(fiber.Map{
			"err": err.Error(),
		})
	}
	data, errorr := user.Get()
	if errorr != nil && errorr == sql.ErrNoRows {
		return c.JSON(fiber.Map{
			"err": errorr.Error(),
		})
	}
	if data != nil {
		return c.JSON(fiber.Map{
			"err": "Account is alreday availabke with this email address.",
		})
	}
	eror := user.Create()
	if eror != nil {
		return c.JSON(fiber.Map{
			"err": eror.Error(),
		})
	}
	return c.SendString("Account created.")
}

func LoginUser(c *fiber.Ctx) error {
	user := new(models.User)
	eror := c.BodyParser(&user)
	if eror != nil {
		return c.JSON(fiber.Map{
			"err": eror.Error(),
		})
	}
	data, err := user.Get()
	if err != nil {
		return c.JSON(fiber.Map{
			"err": err.Error(),
		})
	}
	users := data.(*models.User)
	ok := bcrypt.CompareHashAndPassword([]byte(users.Passowrd), []byte(user.Passowrd))
	if ok != nil {
		return c.JSON(fiber.Map{
			"err": ok.Error(),
		})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1 * time.Hour)
	claims["username"] = users.Username
	claims["email"] = users.Email

	t, tErr := token.SignedString([]byte("mypasswordaaltufaltukhaopio"))
	if tErr != nil {
		return c.JSON(fiber.Map{
			"err": tErr.Error(),
		})
	}
	cookies := fiber.Cookie{
		Name:     "access_token",
		Value:    t,
		MaxAge:   0,
		Expires:  time.Now().Add(48 * time.Hour),
		Secure:   true,
		HTTPOnly: true,
	}
	c.Cookie(&cookies)
	return c.JSON(fiber.Map{
		"username":     users.Username,
		"email":        users.Email,
		"access_token": t,
	})
}
