package utils

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func generateOTPConnectionString() string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(time.Now().String()), 2)
	return string(bytes)
}
