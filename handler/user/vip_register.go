package user

import (
	"context"
	"fmt"
	"log"
	"strings"
	"technexRegistration/config"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const vipTicketName = "TECHNEX VIP CARD"

type vipRegistrationRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Company       string `json:"company"`
	AadhaarNumber string `json:"aadhaarNumber"`
	Phone         string `json:"phone,omitempty"`
}

func RegisterVIP(c *fiber.Ctx) error {
	apiKey := c.Get("api-key")
	adminKey := config.Config("admin_key")

	if apiKey == "" {
		return c.Status(401).JSON(fiber.Map{"message": "api-key header is required"})
	}
	if apiKey != adminKey {
		return c.Status(401).JSON(fiber.Map{"message": "invalid api-key"})
	}

	var payload vipRegistrationRequest
	if err := c.BodyParser(&payload); err != nil {
		log.Printf("VIP registration: invalid request body: %v", err)
		return utils.ResponseMsg(c, 400, "invalid request body", nil)
	}

	payload.Name = strings.TrimSpace(payload.Name)
	payload.Email = normalizeEmail(payload.Email)
	payload.Company = strings.TrimSpace(payload.Company)
	payload.AadhaarNumber = strings.TrimSpace(payload.AadhaarNumber)
	payload.Phone = strings.TrimSpace(payload.Phone)

	if payload.Name == "" || payload.Email == "" || payload.Company == "" {
		return utils.ResponseMsg(c, 400, "name, email and company are required", nil)
	}
	if payload.AadhaarNumber == "" {
		return utils.ResponseMsg(c, 400, "aadhaarNumber is required", nil)
	}

	db, err := database.Connect()
	if err != nil {
		log.Printf("VIP registration: database connection error: %v", err)
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	ctx := context.Background()
	var existing models.Users
	err = db.Collection("users").FindOne(ctx, bson.M{"email": payload.Email}).Decode(&existing)
	if err == nil {
		return c.Status(409).JSON(fiber.Map{"message": "Email already exists"})
	}
	if err != nil && err != mongo.ErrNoDocuments {
		log.Printf("VIP registration: email lookup error: %v", err)
		return c.Status(500).JSON(fiber.Map{"message": "Failed to validate email uniqueness"})
	}

	numCollection := db.Collection("num")
	var currentNum struct {
		Number int `bson:"number"`
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)
	if err := numCollection.FindOneAndUpdate(ctx, bson.M{}, bson.M{"$inc": bson.M{"number": 1}}, opts).Decode(&currentNum); err != nil {
		log.Printf("VIP registration: counter update failed: %v", err)
		return c.Status(500).JSON(fiber.Map{"message": "Failed to generate Technex ID"})
	}

	zeroPadding := 4 - len(fmt.Sprintf("%d", currentNum.Number))
	technexID := "TX26"
	for i := 0; i < zeroPadding; i++ {
		technexID += "0"
	}
	technexID += fmt.Sprintf("%d", currentNum.Number)

	identitySuffix := fmt.Sprintf("%d", time.Now().UnixNano())
	username := fmt.Sprintf("vip-%s", identitySuffix)
	password := utils.HashPassword(fmt.Sprintf("vip-pwd-%s", identitySuffix))
	phone := payload.Phone
	if phone == "" {
		phone = "0000000000"
	}

	user := models.Users{
		Name:             payload.Name,
		Username:         username,
		Password:         password,
		Institute:        payload.Company,
		Company:          payload.Company,
		City:             payload.Company,
		Gender:           "others",
		Year:             0,
		Branch:           "VIP",
		Phone:            phone,
		ReferralCode:     "VIP",
		Email:            payload.Email,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		TechnexID:        technexID,
		RegisteredEvents: []string{},
		Tickets: []models.Ticket{
			{
				Name:          vipTicketName,
				TicketURL:     "",
				InvoiceURL:    "",
				Accommodation: false,
			},
		},
		Accommodation: false,
		EmailVerified: true,
		AadhaarNumber: payload.AadhaarNumber,
	}

	if _, err := db.Collection("users").InsertOne(ctx, user); err != nil {
		log.Printf("VIP registration: insert error: %v", err)
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	qrToken, _ := utils.SerialiseQR(user.TechnexID)
	return c.Status(201).JSON(fiber.Map{
		"technexId":     user.TechnexID,
		"name":          user.Name,
		"email":         user.Email,
		"company":       user.Company,
		"aadhaarNumber": user.AadhaarNumber,
		"qrToken":       qrToken,
	})
}
