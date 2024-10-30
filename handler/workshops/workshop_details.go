package workshops

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
)

func GetWorkshopDetails(c *fiber.Ctx) error {
	workshop := new(models.Workshop)
	var ctx = context.Background()

	db, err := database.Connect()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	if err := c.BodyParser(workshop); err != nil {
		return utils.ResponseMsg(c, 400, err.Error(), nil)
	}

	filter := bson.D{{Key: "name", Value: workshop.Name}}
	var foundWorkshop models.Workshop
	err = db.Collection("workshops").FindOne(ctx, filter).Decode(&foundWorkshop)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(foundWorkshop)
}
