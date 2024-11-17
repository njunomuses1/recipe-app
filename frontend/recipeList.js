import React, { useState, useEffect } from 'react';
import axios from 'axios';

const RecipeList = () => {
  const [recipes, setRecipes] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8080/api/recipes')
      .then(response => {
        setRecipes(response.data);
      })
      .catch(error => {
        console.error("There was an error fetching the recipes!", error);
      });
  }, []);

  return (
    <div>
      <h1>Recipes</h1>
      <ul>
        {recipes.map(recipe => (
          <li key={recipe.id}>
            <h2>{recipe.name}</h2>
            <p>{recipe.instructions}</p>
            <p><strong>Ingredients:</strong> {recipe.ingredients.join(", ")}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default RecipeList;
