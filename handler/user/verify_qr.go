package user

import (
	"github.com/gofiber/fiber/v2"
	"technexRegistration/utils"
)

func VerifyQR(c *fiber.Ctx) error {
	var body struct {
		QRToken string `json:"qr_token"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	username, err := utils.DeserialiseQR(body.QRToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid QR token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"username": username,
	})
}
