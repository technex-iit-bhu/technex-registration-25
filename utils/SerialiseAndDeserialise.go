package utils

import (
	"fmt"
	"technexRegistration/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var serialKey = []byte(config.Config("JWT_SECRET"))
var gmailKey = []byte(config.Config("GMAIL_SECRET"))
var githubKey = []byte(config.Config("GITHUB_SECRET"))
var recoveryKey = []byte(config.Config("RECOVERY_SECRET"))
var qrKey = []byte(config.Config("QR_SECRET"))
var refreshKey = []byte(config.Config("REFRESH_SECRET"))

func SerialiseAccessToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"type":     "access",
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
		// "exp":	  	time.Now().Add(30 * time.Second).Unix(), // 10 seconds for testing
		"iat":      time.Now().Unix(),
	})
	signedToken, err := token.SignedString(serialKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseAccessToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(serialKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	// Check token type
	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "access" {
		return "", fmt.Errorf("invalid token type")
	}

	// Check expiration
	exp, ok := claims["exp"].(float64)
	if !ok {
		return "", fmt.Errorf("invalid expiration")
	}
	if time.Now().Unix() > int64(exp) {
		return "", fmt.Errorf("token expired")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("invalid username claim")
	}

	return username, nil
}

func SerialiseRefreshToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"type":     "refresh",
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(), // 7 days
		// "exp":      time.Now().Add(2 * time.Minute).Unix(),
		"iat":      time.Now().Unix(),
	})
	signedToken, err := token.SignedString(refreshKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseRefreshToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	// Check token type
	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		return "", fmt.Errorf("invalid token type")
	}

	// Check expiration
	exp, ok := claims["exp"].(float64)
	if !ok {
		return "", fmt.Errorf("invalid expiration")
	}
	if time.Now().Unix() > int64(exp) {
		return "", fmt.Errorf("token expired")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("invalid username claim")
	}

	return username, nil
}


func SerialiseUser(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	signedToken, err := token.SignedString(serialKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseUser(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(serialKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims["username"].(string), nil
}

func SerialiseQR(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	signedToken, err := token.SignedString(qrKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseQR(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(qrKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims["username"].(string), nil
}

func SerialiseGmailToken(gmail string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"gmail": gmail,
	})
	signedToken, err := token.SignedString(gmailKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseGmailToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(gmailKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["gmail"].(string), nil
}

func SerialiseGithubToken(github string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"github": github,
	})
	signedToken, err := token.SignedString(githubKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseGithubToken(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(githubKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["github"].(string), nil
}

func SerialiseRecovery(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   username,
		"expires_at": time.Now().Add(10 * time.Minute).Unix(),
	})
	signedToken, err := token.SignedString(recoveryKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func DeserialiseRecovery(signedToken string) (string, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(recoveryKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	if time.Now().After(time.Unix(int64(claims["expires_at"].(float64)), 0)) {
		return "", fmt.Errorf("expired")
	}
	return claims["username"].(string), nil
}
