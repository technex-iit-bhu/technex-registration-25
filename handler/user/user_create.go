package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"time"
)

func CreateUsers(c *fiber.Ctx) error {
	users := new(models.Users)
	var ctx = context.Background()
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()

	if err := c.BodyParser(users); err != nil {
		return utils.ResponseMsg(c, 400, err.Error(), nil)
	} else {
		users.Password = utils.HashPassword(users.Password)
		if r, err := db.Collection("users").InsertOne(ctx, users); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		} else {
			return c.Status(201).JSON(fiber.Map{"id": r.InsertedID})
		}
	}
}
