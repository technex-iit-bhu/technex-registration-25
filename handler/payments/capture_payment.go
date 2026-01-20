package payments

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"technexRegistration/database"
	"technexRegistration/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var allEvents = []string{
	"International Coding Marathon (ICM)",
	"MLWare",
	"Hack It Out",
	"Capture The Flag (CTF)",
	"Game Jam",
	"SkyGlide",
	"AeroVerse",
	"DroneTech",
	"Botstacle Challenge",
	"Micromouse",
	"MazeX",
	"Robowars",
	"Star-Hopping Challenge",
	"AstroQuiz",
	"Stellar Analytics",
	"Solid Boost",
	"AI-Quisition",
	"Prodnosis",
	"Technalytics",
	"Consultathon",
	"NitiVerse",
	"Eco Hackathon",
	"Green Ideathon",
	"Eureka",
	"EngiNX: The Thermal Edition",
	"Axelerate",
	"Robosoccer",
	"Boat Racing Competition",
	"Algozen",
	"Pokermania",
	"CogniQuest",
	"MarketSmith",
}

var allEventTickets = []string{
	"Technex (Event + Accommodation) Card",
	"Technex Events Card",
	"Test all events card",
}

var singleEventTickets = []string{
	"Technex Single Event Card",
	"Technex Single Event + Accomodation Card",
	"Test single event card",
}

type TicketDetails struct {
	TicketName string `json:"Ticket Name"`
}

type AttendeeDetails struct {
	Email      string        `json:"Email address"`
	TechnexId  string        `json:"Technex ID"`
	Event      string        `json:"Event "`
	Ticket     TicketDetails `json:"Ticket Details"`
	TicketURL  string        `json:"Ticket URL"`
	InvoiceURL string        `json:"Invoice URL"`
}

type Details struct {
	AttDetails AttendeeDetails `json:"Attendee Details"`
}

type Body struct {
	Data Details `json:"Data"`
}

func getEventsFromAttendeeDetails(AttDetails AttendeeDetails) []string {
	newItems := []string{}
	if slices.Contains(singleEventTickets, AttDetails.Ticket.TicketName) {
		newItems = []string{AttDetails.Event}
	} else if slices.Contains(allEventTickets, AttDetails.Ticket.TicketName) {
		newItems = allEvents
	}
	return newItems
}

func updateUserEvents(technexId string, newItems []string, ticket models.Ticket) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	update := bson.M{
		"$addToSet": bson.M{
			"registeredEvents": bson.M{
				"$each": newItems,
			},
			"tickets": ticket,
		},
	}

	if ticket.Accommodation {
		update["$set"] = bson.M{"accommodation": true}
	}

	result, err := db.Collection("users").UpdateOne(context.Background(), bson.D{{Key: "technexId", Value: technexId}}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("user does not exist")
	}
	return nil

}

func CapturePayments(c *fiber.Ctx) error {
	var body Body
	c.BodyParser(&body)
	out, _ := json.MarshalIndent(body, "", "  ")
	fmt.Println(string(out))
	newItems := getEventsFromAttendeeDetails(body.Data.AttDetails)

	ticketName := body.Data.AttDetails.Ticket.TicketName
	hasAccommodation := strings.Contains(ticketName, "Accomodation") || strings.Contains(ticketName, "Accommodation")

	ticket := models.Ticket{
		Name:          ticketName,
		TicketURL:     body.Data.AttDetails.TicketURL,
		InvoiceURL:    body.Data.AttDetails.InvoiceURL,
		Accommodation: hasAccommodation,
	}

	err := updateUserEvents(body.Data.AttDetails.TechnexId, newItems, ticket)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "user updated successfully"})
}
