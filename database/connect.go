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

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Check the connection : Send a Ping
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client.Database(config.Config("MONGO_DB_NAME")), nil
}
