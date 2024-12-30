package workshops

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
)

func BulkInsertWorkshops(c *fiber.Ctx) error {
	ctx := context.Background()
	token := c.Get("Authorization")[7:]
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	var workshops []models.Workshop
	if err := c.BodyParser(&workshops); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	var docs []interface{}
	for _, event := range workshops {
		docs = append(docs, event)
	}

	res, err := db.Collection("workshops").InsertMany(ctx, docs)
	if err != nil {
		return utils.ResponseMsg(c, 400, "Failed to insert workshops", nil)
	}

	return c.Status(200).JSON(fiber.Map{"message": "Workshops inserted", "insertedIDs": res.InsertedIDs})
}
