package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection interface {
	Close()
	DB() *mongo.Database
}

type conn struct {
	client *mongo.Client
}

func NewConnection() Connection {
	var c conn
	var err error
	url := getURL()

	clientOptions := options.Client().ApplyURI(url)

	c.client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panicln(err.Error())
	}
	return &c
}

func (c *conn) Close() {
	c.Close()
}

func (c *conn) DB() *mongo.Database {
	return c.client.Database("fiber")
}

func getURL() string {
	return os.Getenv("MONGO_URL")
}

//func ConnDB(databaseURL string) (*mongo.Client, error) {
//	clientOptions := options.Client().ApplyURI(databaseURL)
//	client, err := mongo.Connect(context.Background(), clientOptions)
//	if err != nil {
//		return nil, fmt.Errorf("failed to connect to DB: %w", err)
//	}
//
//	err = client.Ping(context.Background(), nil)
//	if err != nil {
//		return nil, fmt.Errorf("failed pinging DB: %w", err)
//	}
//
//	log.Println("Connected to database successfully!")
//	return client, nil
//}
