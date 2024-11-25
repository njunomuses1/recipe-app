package services

import (
	"context"
	"log"
	"time"

	"github.com/njunomoses1/Recipe/backend/db"
	"github.com/njunomoses1/Recipe/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetRecipes retrieves all recipes from MongoDB with pagination support.
func GetRecipes(page, pageSize int) ([]models.Recipe, error) {
	collection := db.Database.Collection("recipes")
	var recipes []models.Recipe

	// Add context with timeout to avoid hanging
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set options for pagination (skip and limit)
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * pageSize))
	findOptions.SetLimit(int64(pageSize))

	cursor, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Println("Error fetching recipes:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Process the cursor
	for cursor.Next(ctx) {
		var recipe models.Recipe
		if err := cursor.Decode(&recipe); err != nil {
			log.Println("Error decoding recipe:", err)
			continue // skip this record if decoding fails
		}
		recipes = append(recipes, recipe)
	}

	// Check if there were any errors during iteration
	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return recipes, nil
}

// CreateRecipe adds a new recipe to MongoDB and logs any errors.
func CreateRecipe(recipe *models.Recipe) (*mongo.InsertOneResult, error) {
	collection := db.Database.Collection("recipes")

	// Add context with timeout to avoid hanging
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, recipe)
	if err != nil {
		log.Println("Error inserting recipe:", err)
		return nil, err
	}

	return result, nil
}
