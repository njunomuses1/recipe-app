package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// Assuming db package is responsible for MongoDB initialization
	// Assuming routes package handles route definitions
)

var client *mongo.Client
var recipeCollection *mongo.Collection

// Recipe struct represents the structure of a recipe in the MongoDB database.
type Recipe struct {
	ID           string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string   `json:"name" bson:"name"`
	Ingredients  []string `json:"ingredients" bson:"ingredients"`
	Instructions string   `json:"instructions" bson:"instructions"`
}

// Initialize MongoDB connection
func initMongoDB() {
	// Connect to MongoDB (replace with your connection string)
	clientOptions := options.Client().ApplyURI("mongodb+srv://your-mongo-uri")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the collection for recipes
	recipeCollection = client.Database("recipe_db").Collection("recipes")
	log.Println("Connected to MongoDB")
}

// Load all recipes from MongoDB
func loadRecipes() ([]Recipe, error) {
	var recipes []Recipe
	cursor, err := recipeCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var recipe Recipe
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

// Split ingredients query into a slice of strings
func splitIngredients(input string) []string {
	// Split by comma and trim spaces
	ingredients := strings.Split(input, ",")
	for i, ingredient := range ingredients {
		ingredients[i] = strings.TrimSpace(ingredient)
	}
	return ingredients
}

// Check if all ingredients are present in the recipe
func containsAllIngredients(recipeIngredients, availableIngredients []string) bool {
	for _, reqIngredient := range recipeIngredients {
		found := false
		for _, availableIngredient := range availableIngredients {
			if reqIngredient == availableIngredient {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Get all recipes that match the given ingredients
func getRecipes(c *gin.Context) {
	ingredientsParam := c.DefaultQuery("ingredients", "")
	ingredients := splitIngredients(ingredientsParam)

	recipes, err := loadRecipes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load recipes"})
		return
	}

	var matchedRecipes []Recipe
	for _, recipe := range recipes {
		if containsAllIngredients(recipe.Ingredients, ingredients) {
			matchedRecipes = append(matchedRecipes, recipe)
		}
	}

	c.JSON(http.StatusOK, matchedRecipes)
}

// Create a new recipe in MongoDB
func createRecipe(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe data"})
		return
	}

	// Insert the recipe into MongoDB
	result, err := recipeCollection.InsertOne(context.Background(), recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert recipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result.InsertedID})
}

func main() {
	// Initialize MongoDB connection
	initMongoDB()

	// Set up Gin routes
	r := gin.Default()

	// Enable CORS to allow requests from the frontend (localhost:3000)
	r.Use(cors.Default())

	// Define routes
	r.GET("/api/recipes", getRecipes)
	r.POST("/api/recipes", createRecipe)

	// Start the server
	r.Run(":8080") // Port 8080
}
