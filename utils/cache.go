package utils

import (
	"context"
	"encoding/json"
	"time"

	"technexRegistration/database"
	"technexRegistration/models"
)

const (
	UserProfileTTL       = 60 * time.Minute
	userProfileKeyPrefix = "user_profile:"
)

// SetUserProfile stores the user's profile in Redis with a TTL.
// It fails silently if Redis is not configured or an error occurs,
// to ensure the main application flow is not interrupted.
func SetUserProfile(username string, u models.Users) {
	rdb := database.GetRedis()
	if rdb == nil {
		return
	}

	data, err := json.Marshal(u)
	if err != nil {
		// In a production environment, you might want to log this error
		return
	}

	// Set the key with the specified TTL
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = rdb.Set(ctx, userProfileKeyPrefix+username, data, UserProfileTTL).Err()
}

// GetUserProfile retrieves a user's profile from Redis.
// Returns (value, true) on hit, (empty, false) on miss or error.
func GetUserProfile(username string) (models.Users, bool) {
	rdb := database.GetRedis()
	if rdb == nil {
		return models.Users{}, false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := rdb.Get(ctx, userProfileKeyPrefix+username).Bytes()
	if err != nil {
		return models.Users{}, false
	}

	var u models.Users
	if err := json.Unmarshal(val, &u); err != nil {
		return models.Users{}, false
	}

	return u, true
}

// DeleteUserProfile invalidates the cached profile for a user.
func DeleteUserProfile(username string) {
	rdb := database.GetRedis()
	if rdb == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = rdb.Del(ctx, userProfileKeyPrefix+username).Err()
}
