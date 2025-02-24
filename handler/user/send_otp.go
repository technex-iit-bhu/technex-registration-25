package user

import (
    "context"
    "fmt"
    "math/rand"
    "net/smtp"
    "time"
    "technexRegistration/config"
    "technexRegistration/database"
    "technexRegistration/models"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func SendOTP(c *fiber.Ctx) error {
    var body struct {
        Email string `json:"email"`
    }

    if err := c.BodyParser(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    db, err := database.Connect()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
    }

    // Check if user exists
    var user models.Users
    err = db.Collection("users").FindOne(
        context.Background(),
        bson.D{{Key: "email", Value: body.Email}},
    ).Decode(&user)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user does not exist"})
    }

    // Generate a 6-digit OTP
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))
    otpCode := rng.Intn(900000) + 100000 // range 100000-999999

    // Send OTP via email
    smtpHost := config.Config("SMTP_HOST")
    smtpPort := config.Config("SMTP_PORT")
    senderEmail := config.Config("SMTP_EMAIL")
    senderPassword := config.Config("SMTP_PASSWORD")

    auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
    to := []string{body.Email}
    subject := "Subject: Password Reset OTP\n"
    msgBody := fmt.Sprintf("Your OTP for password reset is: %d", otpCode)
    message := []byte(subject + "\n" + msgBody)

    err = smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, to, message)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to send OTP"})
    }

    otpDoc := models.Otp{
        ID:        primitive.NewObjectID(),
        UserID:    user.ID,
        Code:      otpCode,
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().Add(10 * time.Minute),
        Used:      false,
    }

    _, err = db.Collection("otps").InsertOne(context.Background(), otpDoc)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to store OTP",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "OTP sent successfully",
    })
}
