package services

import (
	"context"
	"github.com/njunomoses1/Recipe/backend/db"
	"github.com/njunomoses1/Recipe/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// GetRecipes retrieves all recipes from MongoDB
func GetRecipes() ([]models.Recipe, error) {
	collection := db.Database.Collection("recipes")
	var recipes []models.Recipe

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var recipe models.Recipe
		if err := cursor.Decode(&recipe); err != nil {
			log.Println("Error decoding recipe:", err)
			continue
		}
		recipes = append(recipes, recipe)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

// CreateRecipe adds a new recipe to MongoDB
func CreateRecipe(recipe *models.Recipe) (*mongo.InsertOneResult, error) {
	collection := db.Database.Collection("recipes")
	return collection.InsertOne(context.Background(), recipe)
}
