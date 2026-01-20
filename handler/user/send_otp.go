package user

import (
    "context"
    "fmt"
    "math/rand"
    "time"
    "technexRegistration/config"
    "technexRegistration/database"
    "technexRegistration/models"

    "github.com/gofiber/fiber/v2"
    // "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/resend/resend-go"
    "strconv"
)
func SendOTP(c *fiber.Ctx) error {
    var body struct {
        Email string `json:"email"`
    }

    if err := c.BodyParser(&body); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    db, err := database.Connect()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
    }
    ctx := context.Background()

    // Check if user exists
    var user models.Users
    err = db.Collection("users").FindOne(ctx, bson.M{
        "email": body.Email,
    }).Decode(&user)

    if err != nil {
        return c.Status(404).JSON(fiber.Map{"message": "User not found"})
    }

    // Generate OTP as string
    otpCode := fmt.Sprintf("%06d", rand.Intn(1000000))

    // Delete previous OTPs
    db.Collection("otps").DeleteMany(ctx, bson.M{
        "user_id": user.ID,
        "used": false,
    })

    // Send OTP via Resend
    resendKey := config.Config("RESEND_API_KEY")
    fromEmail := config.Config("EMAIL_FROM")

    client := resend.NewClient(resendKey)

    _, err = client.Emails.Send(&resend.SendEmailRequest{
        From:    fromEmail,
        To:      []string{body.Email},
        Subject: "Your OTP for password reset",
        Html:    fmt.Sprintf("<h2>Your OTP is: %s</h2><p>Valid for 10 minutes</p>", otpCode),
    })

    if err != nil {
        fmt.Println("RESEND ERROR:", err)
        return c.Status(500).JSON(fiber.Map{"message": "Failed to send OTP"})
    }

    otpDoc := models.Otp{
        ID:        primitive.NewObjectID(),
        UserID:    user.ID,
        Code:      func() int { v, _ := strconv.Atoi(otpCode); return v }(),
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().Add(10 * time.Minute),
        Used:      false,
    }

    _, err = db.Collection("otps").InsertOne(ctx, otpDoc)
    if err != nil {
        fmt.Println("OTP DB ERROR:", err)
        return c.Status(500).JSON(fiber.Map{"message": "Failed to save OTP"})
    }

    return c.JSON(fiber.Map{"message": "OTP sent successfully"})
}
