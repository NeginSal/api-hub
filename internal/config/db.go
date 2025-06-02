package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found. Using default values")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}

	MongoClient = client
	log.Println("Connected to MongoDB Atlas")
}

func GetCollection(dbName, colName string) *mongo.Collection {
	return MongoClient.Database(dbName).Collection(colName)
}
