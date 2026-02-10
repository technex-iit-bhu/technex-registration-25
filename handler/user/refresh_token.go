package user

import (
	"github.com/gofiber/fiber/v2"
	"technexRegistration/utils"
)

func RefreshToken(c *fiber.Ctx) error {

	refreshToken := c.Cookies("refresh_token")
	
	if refreshToken == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "refresh token not found",
		})
	}

	username, err := utils.DeserialiseRefreshToken(refreshToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "invalid or expired refresh token",
		})
	}

	newAccessToken, err := utils.SerialiseAccessToken(username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to generate access token",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"access_token": newAccessToken,
		"token_type":   "Bearer",
		"expires_in":   7200, // 2 hours
		// "expires_in":   30, //testing
	})
}