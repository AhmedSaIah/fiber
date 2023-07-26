package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBCollections struct {
	UserCollection *mongo.Collection
}

var DB *DBCollections

func ConnDB(databaseURL string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(databaseURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed pinging DB: %w", err)
	}

	log.Println("Connected to database successfully!")
	return client, nil
}

func getCollection(collection string) *mongo.Collection {
	m, err := ConnDB(os.Getenv("MONGO_URL"))
	if err != nil {
		panic(err)
	}
	coll := m.Database("fiber").Collection(collection)
	return coll
}

func InitDB(userCollection *mongo.Collection) {
	DB = &DBCollections{
		UserCollection: userCollection,
	}
}
func CreateCollections() {
	m, err := ConnDB(os.Getenv("MONGO_URL"))
	if err != nil {
		panic(err)
	}
	collectionUsers := m.Database("fiber").Collection("users")
	InitDB(collectionUsers)
}
