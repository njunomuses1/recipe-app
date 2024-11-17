import React, { useState } from 'react';
import axios from 'axios';

function App() {
  // State variables
  const [ingredients, setIngredients] = useState('');
  const [recipes, setRecipes] = useState([]);
  const [error, setError] = useState('');

  // Handle user input
  const handleIngredientsChange = (e) => {
    setIngredients(e.target.value);
  };

  // Fetch recipes from the backend
  const getRecipes = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/api/recipes?ingredients=${ingredients}`);
      setRecipes(response.data);
      setError('');
    } catch (err) {
      setError('Error fetching recipes');
    }
  };

  return (
    <div className="App">
      <h1>Dynamic Recipe Generator</h1>

      <div>
        <input
          type="text"
          value={ingredients}
          onChange={handleIngredientsChange}
          placeholder="Enter ingredients (comma separated)"
        />
        <button onClick={getRecipes}>Find Recipes</button>
      </div>

      {error && <p style={{ color: 'red' }}>{error}</p>}

      {recipes.length > 0 && (
        <div>
          <h2>Recipes:</h2>
          <ul>
            {recipes.map((recipe, index) => (
              <li key={index}>
                <h3>{recipe.name}</h3>
                <p><strong>Ingredients:</strong> {recipe.ingredients.join(', ')}</p>
                <p><strong>Instructions:</strong> {recipe.instructions}</p>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}

export default App;
