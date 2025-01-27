package payments

import (
	"context"
	"slices"
	"technexRegistration/database"
	"technexRegistration/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var allEvents = []string{
	"DroneTech",
	"AeroVerse",
	"SkyGlide",
	"International Coding Marathon",
	"Capture the Flag",
	"Hack it out",
	"1 Billion row challenge",
	"Axelerate",
	"Cadastrophe",
	"Robo Soccer",
	"SIMUSOLVE CHALLENGE",
	"Fake investors",
	"Climate buzzer",
	"Eureka",
	"Eco hackathon",
	"AlgoZen",
	"CogniQuest",
	"Pokermania",
	"CryptoRush",
	"IOmatic",
	"Soft-corner",
	"Terravate",
	"Consultathon",
	"Prodonosis",
	"Technalatics",
	"Capital Quest",
	"Robowars",
	"Micromouse",
	"Botstacle Challenge",
	"Mazex",
	"Scientists of Utopia",
	"Solid-Boost",
	"Stellar Analytics",
	"Astro-Quiz",
}

var allEventTickets = []string{
	"Technex Early Bird Event Card",
	"Technex Early Bird (Event + Food) Card",
	"Test all events card",
	"Technex Events Card",
}

var singleEventTickets = []string{
	"Technex Single Event Card",
	"Technex Single Event + Accomodation Card",
	"Test single event card",
	"Technex (Event + Accommodation) Card",
}

type TicketDetails struct {
	TicketName string `json:"Ticket Name"`
}

type AttendeeDetails struct {
	Email     string        `json:"Email address"`
	TechnexId string        `json:"Technex ID"`
	Event     string        `json:"Event "`
	Ticket    TicketDetails `json:"Ticket Details"`
}

type Details struct {
	AttDetails AttendeeDetails `json:"Attendee Details"`
}

type Body struct {
	Data Details `json:"Data"`
}

func CapturePayments(c *fiber.Ctx) error {
	var body Body
	c.BodyParser(&body)

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	var user models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "technexId", Value: body.Data.AttDetails.TechnexId}}).Decode(&user)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}

	newItems := []string{}

	if slices.Contains(singleEventTickets, body.Data.AttDetails.Ticket.TicketName) {
		newItems = []string{body.Data.AttDetails.Event}
	} else if slices.Contains(allEventTickets, body.Data.AttDetails.Ticket.TicketName) {
		newItems = allEvents
	}

	result, _ := db.Collection("users").UpdateOne(context.Background(), bson.D{{Key: "technexId", Value: body.Data.AttDetails.TechnexId}},
		bson.M{
			"$addToSet": bson.M{
				"registeredEvents": bson.M{
					"$each": newItems,
				},
			}})
	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "user updated successfully"})
}
