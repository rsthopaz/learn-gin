package db

import (
	"context"
	"fmt"
	"os"
	"time"
	"log"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	err := godotenv.Load("../../.env")
	if err != nil {
    log.Println("No .env file found")
}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	if mongoURI == "" || dbName == "" {
		log.Fatal("DB Credentails are missing")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed connect to MongoDB:", err)
	}

	err = client.Ping(ctx,nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	fmt.Println("Connected to MongoDB")
	DB = client.Database(dbName)

}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}