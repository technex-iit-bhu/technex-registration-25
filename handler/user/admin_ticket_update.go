package user

import (
	"context"
	"strings"
	"technexRegistration/config"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type AdminTicketUpdateBody struct {
	TicketName    string `json:"ticketName"`
	TicketURL     string `json:"ticketUrl"`
	InvoiceURL    string `json:"invoiceUrl"`
	Accommodation *bool  `json:"accommodation,omitempty"`
}

func AdminUpsertTicketByTechnexID(c *fiber.Ctx) error {
	apiKey := strings.TrimSpace(c.Get("api-key"))
	adminKey := strings.TrimSpace(config.Config("admin_key"))
	if apiKey == "" || adminKey == "" || apiKey != adminKey {
		return c.Status(401).JSON(fiber.Map{"message": "invalid api key"})
	}

	technexID := strings.TrimSpace(c.Params("technexId"))
	if technexID == "" {
		return c.Status(400).JSON(fiber.Map{"message": "technexId is required"})
	}

	var body AdminTicketUpdateBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request body"})
	}

	body.TicketName = strings.TrimSpace(body.TicketName)
	if body.TicketName == "" {
		return c.Status(400).JSON(fiber.Map{"message": "ticketName is required"})
	}

	accommodation := false
	if body.Accommodation != nil {
		accommodation = *body.Accommodation
	} else {
		normalized := strings.ToLower(body.TicketName)
		accommodation = strings.Contains(normalized, "accomodation") || strings.Contains(normalized, "accommodation")
	}

	ticket := models.Ticket{
		Name:          body.TicketName,
		TicketURL:     body.TicketURL,
		InvoiceURL:    body.InvoiceURL,
		Accommodation: accommodation,
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	filter := bson.M{"technexId": technexID}
	updateSet := bson.M{"updatedAt": time.Now()}
	if accommodation {
		updateSet["accommodation"] = true
	}

	result, err := db.Collection("users").UpdateOne(
		context.Background(),
		filter,
		bson.M{
			"$addToSet": bson.M{"tickets": ticket},
			"$set":      updateSet,
		},
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}

	var cachedUser struct {
		Username string `bson:"username"`
	}
	_ = db.Collection("users").FindOne(context.Background(), filter).Decode(&cachedUser)
	if cachedUser.Username != "" {
		utils.DeleteUserProfile(cachedUser.Username)
	}

	return c.Status(200).JSON(fiber.Map{
		"message":   "ticket updated successfully",
		"technexId": technexID,
		"ticket":    ticket,
	})
}
