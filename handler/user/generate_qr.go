package user

import (
	"technexRegistration/utils"

	"github.com/gofiber/fiber/v2"
)

func GenerateQR(c *fiber.Ctx) error {
	var qrBody struct {
		Username string `json:"username" bson:"username"`
		Name string `json:"name" bson:"name"`
	}

	if err := c.BodyParser(&qrBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request qrBody",
		})
	}

	qrToken, err := utils.SerialiseQR(qrBody.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate QR token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"qr_token": qrToken,
	})
}
