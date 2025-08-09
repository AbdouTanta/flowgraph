package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// init makes sure the MongoDB client is initialized when the package is imported.
func InitMongoDbClient() *mongo.Database {
	// Get the MongoDB URI from the environment variable
	var uri string
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
	}

	// Create a new client and connect to the server
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	//Set the global Database variable to the "flowgraph" database
	database := client.Database("flowgraph")
	// Ping the database to verify the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// If we reach here, the connection was successful
	fmt.Println("Successfully connected to MongoDB!")
	return database
}
