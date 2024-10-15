package handler

import (
	"github.com/gofiber/fiber/v2"
	"technexRegistration/helpers"
)

func Hello(c *fiber.Ctx) error {
	return helpers.ResponseMsg(c, 200, "Api is running", nil)
}
