package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database




// Initialize MongoDB connection
func InitMongoDB() {
	// Replace with your MongoDB URI (local or Atlas)
	clientOptions := options.Client().ApplyURI("mongodb+srv://mosesnjuno7:due5Ct1p0FQU2wbh@recipe.ok1di.mongodb.net/Recipe?retryWrites=true&w=majority") // Use your URI here

	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping the database to confirm connection
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	// Use the database named "recipe_db"
	Database = Client.Database("recipe_db")
	log.Println("Connected to MongoDB")
}
