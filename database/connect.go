package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"technexRegistration/config"
)

func Connect() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(config.Config("MONGO_URI"))

	var ctx = context.Background()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return client.Database(config.Config("MONGO_DB_NAME")), nil
}
