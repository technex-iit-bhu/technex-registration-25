package user

import (
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // Delete cookie
		HTTPOnly: true,
		Secure:   true, // Set to true in production with HTTPS
		SameSite: "Lax",
	})

	return c.Status(200).JSON(fiber.Map{
		"message": "logged out successfully",
	})
}