// package utils

// import (
// 	"fmt"
// 	smtp "net/smtp"
// 	"technexRegistration/config"
// )

// func RecoveryMail(to, token string) error {
// 	password := config.Config("SMTP_PASSWORD")
// 	from := config.Config("SMTP_EMAIL")
// 	smtpHost := config.Config("SMTP_HOST")
// 	smtpPort := config.Config("SMTP_PORT")
// 	message := []byte(fmt.Sprintf("Subject: Recovery\r\n\r\nlink :  https://[FRONTEND_URL]/recovery?token=%s\r\n", token))
// 	auth := smtp.PlainAuth("", from, password, smtpHost)
// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Email Sent Successfully")
// 	return nil
// }

package utils

import (
	"fmt"
	"log"

	"technexRegistration/config"
	"github.com/resend/resend-go"
)

func RecoveryMail(to string, token string) error {
	apiKey := config.Config("RESEND_API_KEY")
	frontendURL := config.Config("FRONTEND_URL")

	from := "Technex <" + config.Config("SENDER_EMAIL") + ">"

	client := resend.NewClient(apiKey)

	recoveryLink := fmt.Sprintf("%s/recovery?token=%s", frontendURL, token)

	htmlContent := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<body style="font-family: Arial, sans-serif;">
		<h2>Password Recovery</h2>
		<p>You requested to reset your password.</p>
		<p>
			<a href="%s" style="
				padding:12px 20px;
				background:#4f46e5;
				color:#ffffff;
				text-decoration:none;
				border-radius:6px;
				display:inline-block;
			">Recover Account</a>
		</p>
		<p>If you didn’t request this, ignore this email.</p>
		<br/>
		<p>— Technex Team</p>
	</body>
	</html>
	`, recoveryLink)

	params := &resend.SendEmailRequest{
		From:    from,
		To:      []string{to},
		Subject: "Password Recovery – Technex",
		Html:    htmlContent,
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		log.Println("Resend error:", err)
		return err
	}

	fmt.Println("Recovery email sent successfully")
	return nil
}
