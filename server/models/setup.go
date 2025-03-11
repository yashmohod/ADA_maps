package models

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const db = "ADAmap"

var client *mongo.Client

func ConnectDatabase() {
	godotenv.Load()

	connectionString := os.Getenv("DB_CONNECTION_STRING")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOption := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	Client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	client = Client
}

func DisconnectDatabase() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
