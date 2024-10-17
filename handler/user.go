package handler

import (
	"context"
	"log"
	"technexRegistration/database"
	"technexRegistration/helpers"
	"technexRegistration/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return helpers.ResponseMsg(c, 200, "Api is running", nil)
}

func CreateUsers(c *fiber.Ctx) error {
	users := new(models.Users)
	var ctx = context.Background()
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err.Error())
		return c.Status(400).JSON(fiber.Map{"message":err.Error()})
	}

	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()
	
	if err := c.BodyParser(users); err != nil {
		return helpers.ResponseMsg(c, 400, err.Error(), nil)
	} else {
		users.Password=helpers.HashPassword(users.Password)
		if r, err := db.Collection("users").InsertOne(ctx, users); err != nil {
			return c.Status(500).JSON(fiber.Map{"message":err.Error()})
		} else {
			return c.Status(201).JSON(fiber.Map{"id":r.InsertedID})
		}
	}
}
