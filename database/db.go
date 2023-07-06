package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnDB(databaseURL string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(databaseURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed pinging DB: %w", err)
	}

	log.Println("Connected to database successfully!")
	return client, nil
}
