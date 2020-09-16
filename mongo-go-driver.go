package connection

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClient() *mongo.Client {
	err := godotenv.Load(".env.local")
	if err != nil {
		godotenv.Load(".env")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	var client *mongo.Client
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + host + ":" + port + "")

	// Client to MongoDB
	client, err = mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func MongoDB() *mongo.Database {
	err := godotenv.Load(".env.local")
	if err != nil {
		godotenv.Load(".env")
	}

	con := MongoClient()
	database := os.Getenv("DB_NAME")
	db := con.Database(database)

	return db
}
