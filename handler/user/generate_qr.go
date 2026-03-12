package user

import (
	"technexRegistration/utils"

	"github.com/gofiber/fiber/v2"
)

func GenerateQ(c *fiber.Ctx) error {
	var qrBody struct {
		TechnexID string `json:"technexId" bson:"technexId"`
		Name      string `json:"name" bson:"name"`
	}

	if err := c.BodyParser(&qrBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request qrBody",
		})
	}

	if qrBody.TechnexID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "technexId required",
		})
	}

	qrToken, err := utils.SerialiseQR(qrBody.TechnexID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate QR token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"qr_token": qrToken,
	})
}
