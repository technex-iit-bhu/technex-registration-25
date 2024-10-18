package handler

import (
	"github.com/gofiber/fiber/v2"
	"technexRegistration/utils"
)

func Hello(c *fiber.Ctx) error {
	return utils.ResponseMsg(c, 200, "Api is running", nil)
}
