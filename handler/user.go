package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"technexRegistration/database"
	"technexRegistration/helpers"
	"technexRegistration/models"
	"time"
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
	}

	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()

	if err := c.BodyParser(users); err != nil {
		return helpers.ResponseMsg(c, 400, err.Error(), nil)
	} else {
		if r, err := db.Collection("users").InsertOne(ctx, users); err != nil {
			return helpers.ResponseMsg(c, 500, "Inserted data unsuccesfully", err.Error())
		} else {
			return helpers.ResponseMsg(c, 200, "Inserted data succesfully", r)
		}
	}
}
