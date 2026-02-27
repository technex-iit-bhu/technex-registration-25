package user

import (
	"fmt"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// normalizeEmail returns a trimmed, lowercase email for consistent lookups.
func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// normalizeUsername removes leading/trailing whitespace but preserves case for display.
func normalizeUsername(username string) string {
	return strings.TrimSpace(username)
}

// usernameCaseInsensitiveFilter builds a case-insensitive Mongo filter for a username.
func usernameCaseInsensitiveFilter(username string) bson.M {
	clean := normalizeUsername(username)
	if clean == "" {
		return bson.M{}
	}
	pattern := fmt.Sprintf("^%s$", regexp.QuoteMeta(clean))
	return bson.M{
		"username": bson.M{
			"$regex":   pattern,
			"$options": "i",
		},
	}
}
