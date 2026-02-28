package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strconv"
	"strings"
	"technexRegistration/config"
)

func sendSMTPHTML(to, subject, htmlBody string) error {
	smtpHost := config.Config("SMTP_HOST")
	smtpPortStr := config.Config("SMTP_PORT")
	smtpEmail := config.Config("SMTP_EMAIL")
	smtpPass := config.Config("SMTP_PASSWORD")

	if smtpHost == "" || smtpPortStr == "" || smtpEmail == "" || smtpPass == "" {
		return fmt.Errorf("missing SMTP config: ensure SMTP_HOST, SMTP_PORT, SMTP_EMAIL, SMTP_PASSWORD are set")
	}

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return fmt.Errorf("invalid SMTP_PORT: %w", err)
	}

	// From header: include a friendly name
	fromHeader := fmt.Sprintf("Technex <%s>", smtpEmail)

	// Basic email headers + HTML MIME
	headers := map[string]string{
		"From":         fromHeader,
		"To":           to,
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": `text/html; charset="UTF-8"`,
	}

	var msg strings.Builder
	for k, v := range headers {
		msg.WriteString(k)
		msg.WriteString(": ")
		msg.WriteString(v)
		msg.WriteString("\r\n")
	}
	msg.WriteString("\r\n")
	msg.WriteString(htmlBody)

	addr := net.JoinHostPort(smtpHost, strconv.Itoa(smtpPort))

	// net/smtp supports STARTTLS via smtp.Client
	c, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("smtp dial: %w", err)
	}
	defer c.Close()

	// STARTTLS (for 587)
	tlsConfig := &tls.Config{
		ServerName: smtpHost,
		MinVersion: tls.VersionTLS12,
	}

	// If server supports STARTTLS, upgrade the connection
	if ok, _ := c.Extension("STARTTLS"); ok {
		if err := c.StartTLS(tlsConfig); err != nil {
			return fmt.Errorf("starttls: %w", err)
		}
	} else {
		// For Gmail: STARTTLS is expected on 587
		return fmt.Errorf("smtp server does not support STARTTLS")
	}

	// Authenticate (only after TLS)
	auth := smtp.PlainAuth("", smtpEmail, smtpPass, smtpHost)
	if ok, _ := c.Extension("AUTH"); ok {
		if err := c.Auth(auth); err != nil {
			return fmt.Errorf("smtp auth: %w", err)
		}
	}

	// Set the envelope
	if err := c.Mail(smtpEmail); err != nil {
		return fmt.Errorf("smtp mail from: %w", err)
	}
	if err := c.Rcpt(to); err != nil {
		return fmt.Errorf("smtp rcpt to: %w", err)
	}

	// Send data
	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("smtp data: %w", err)
	}
	_, err = w.Write([]byte(msg.String()))
	if err != nil {
		_ = w.Close()
		return fmt.Errorf("smtp write: %w", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("smtp close data: %w", err)
	}

	// Quit politely
	if err := c.Quit(); err != nil {
		// not fatal, but useful to log
		log.Println("smtp quit error:", err)
	}

	return nil
}

func RecoveryMail(to string, username string, otp string) error {
	htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<body style="font-family: Arial, sans-serif; background:#0b0b0b; padding:20px;">
 <div style="max-width:480px;margin:auto;background:#000;border:1px solid #f5c54233;border-radius:12px;padding:24px;color:#fff;">
  <h2 style="color:#f5c542;text-align:center;">Password recovery</h2>
  <p>Hello %s,</p>
  <p>Use the OTP below to recover your Technex account:</p>
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

	return sendSMTPHTML(to, "Password Recovery – Technex", htmlContent)
}

func VerificationMail(to string, username string, otp string) error {
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

	return sendSMTPHTML(to, "Verify your Technex email", htmlContent)
}