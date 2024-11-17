package routes

import (
	"github.com/njunomoses1/Recipe/backend/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRecipeRoutes sets up the API routes for recipes
func SetupRecipeRoutes(router *gin.Engine) {
	recipeRoutes := router.Group("/api/recipes")
	{
		recipeRoutes.GET("/", handlers.GetAllRecipes)
		recipeRoutes.POST("/", handlers.CreateRecipe)
	}
}
