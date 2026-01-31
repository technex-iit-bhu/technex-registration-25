package utils

import (
	"fmt"
	"log"

	"technexRegistration/config"
	"github.com/resend/resend-go"
)

func RecoveryMail(to string, username string, otp string) error {
 apiKey := config.Config("RESEND_API_KEY")


 from := "Technex <" + config.Config("SENDER_EMAIL") + ">"


 client := resend.NewClient(apiKey)



 htmlContent := fmt.Sprintf(`

 <!DOCTYPE html>

 <html>

 <body style="font-family: Arial, sans-serif; background:#0b0b0b; padding:20px;">

  <div style="max-width:480px;margin:auto;background:#000;border:1px solid #f5c54233;border-radius:12px;padding:24px;color:#fff;">


   <h2 style="color:#f5c542;text-align:center;">Password recovery</h2>


   <p>Hello %s,</p>


   <p>Use the OTP below to verify your Technex account:</p>


   <div style="text-align:center;margin:30px 0;">

    <span style="font-size:28px;letter-spacing:6px;color:#f5c542;font-weight:bold;">%s</span>

   </div>


   <p>This OTP is valid for 10 minutes.</p>


   <p>If you didn’t request this, please ignore.</p>


   <br/>

   <p style="color:#999;">— Technex Team</p>


  </div>

 </body>

 </html>

 `, username, otp)


 params := &resend.SendEmailRequest{

  From: from,

  To: []string{to},

  Subject: "Password Recovery – Technex",

  Html: htmlContent,

 }


	_, err := client.Emails.Send(params)
		if err != nil {
			log.Println("Recovery mail error:", err)
			return err
		}

	return nil

}

func VerificationMail(to string, username string, otp string) error {

 apiKey := config.Config("RESEND_API_KEY")


 from := "Technex <" + config.Config("SENDER_EMAIL") + ">"


 client := resend.NewClient(apiKey)


 htmlContent := fmt.Sprintf(`

 <!DOCTYPE html>

 <html>

 <body style="font-family: Arial, sans-serif; background:#0b0b0b; padding:20px;">

  <div style="max-width:480px;margin:auto;background:#000;border:1px solid #f5c54233;border-radius:12px;padding:24px;color:#fff;">


   <h2 style="color:#f5c542;text-align:center;">Email Verification</h2>


   <p>Hello %s,</p>


   <p>Use the OTP below to verify your Technex account:</p>


   <div style="text-align:center;margin:30px 0;">

    <span style="font-size:28px;letter-spacing:6px;color:#f5c542;font-weight:bold;">%s</span>

   </div>


   <p>This OTP is valid for 10 minutes.</p>


   <p>If you didn’t request this, please ignore.</p>


   <br/>

   <p style="color:#999;">— Technex Team</p>


  </div>

 </body>

 </html>

 `, username, otp)


 params := &resend.SendEmailRequest{

  From: from,

  To: []string{to},

  Subject: "Verify your Technex email",

  Html: htmlContent,

 }


 _, err := client.Emails.Send(params)

 if err != nil {

  log.Println("Resend verification error:", err)

  return err

 }


 fmt.Println("Verification email sent")

 return nil

}