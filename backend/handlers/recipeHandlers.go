package handlers

import (
	"github.com/njunomoses1/Recipe/backend/models"
	"event-manager/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

// GetAllRecipes retrieves all recipes
func GetAllRecipes(c *gin.Context) {
	recipes, err := services.GetRecipes()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recipes"})
		return
	}

	c.JSON(http.StatusOK, recipes)
}

// CreateRecipe creates a new recipe
func CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe data"})
		return
	}

	result, err := services.CreateRecipe(&recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create recipe"})
		return
	}

	c.JSON(http.StatusCreated, result)
}
