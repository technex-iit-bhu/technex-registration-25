package database

import (
	"context"
	"fmt"
	"technexRegistration/config"

	redis "github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = nil
var connected bool = false
var rdb *redis.Client = nil

func Init() error {
	clientOptions := options.Client().ApplyURI(config.Config("MONGO_URI"))
	var ctx = context.Background()
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	client = c
	connected = true
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// Initialize Redis if URL or address provided. Prefer REDIS_URL.
	redisURL := config.Config("REDIS_URL")
	if redisURL != "" {
		opts, err := redis.ParseURL(redisURL)
		if err != nil {
			return err
		}
		rdb = redis.NewClient(opts)
		if err := rdb.Ping(ctx).Err(); err != nil {
			return err
		}
		fmt.Println("Connected to Redis via REDIS_URL")
	} else {
		fmt.Println("No Redis URL Found")
	}
	return nil
}

func Connect() (*mongo.Database, error) {
	if !connected {
		return nil, fmt.Errorf("client not connected")
	}

	return client.Database(config.Config("MONGO_DB_NAME")), nil
}

// GetClient exposes the MongoDB client for cleanup scenarios.
func GetClient() *mongo.Client {
	return client
}

func Disconnect() error {
	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}
	fmt.Println("connection go brrrr")
	if rdb != nil {
		_ = rdb.Close()
	}
	return nil
}

// GetRedis returns the Redis client (may be nil if not configured)
func GetRedis() *redis.Client {
	return rdb
}
