package user

import (
	"technexRegistration/utils"

	"github.com/gofiber/fiber/v2"
)

func GenerateQR(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	qrToken, err := utils.SerialiseQR(body.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate QR token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"qr_token": qrToken,
	})
}
