package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Backend is Running",
	})
}

func HelloAPI(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "API is running",
	})
}
