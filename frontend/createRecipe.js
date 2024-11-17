import React, { useState } from 'react';
import axios from 'axios';

const CreateRecipe = () => {
  const [name, setName] = useState('');
  const [ingredients, setIngredients] = useState('');
  const [instructions, setInstructions] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();

    const newRecipe = {
      name,
      ingredients: ingredients.split(','),
      instructions,
    };

    axios.post('http://localhost:8080/api/recipes', newRecipe)
      .then(response => {
        console.log("Recipe created:", response.data);
        // Optionally reset the form or show a success message
      })
      .catch(error => {
        console.error("There was an error creating the recipe!", error);
      });
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Name:</label>
        <input type="text" value={name} onChange={(e) => setName(e.target.value)} required />
      </div>
      <div>
        <label>Ingredients (comma separated):</label>
        <input type="text" value={ingredients} onChange={(e) => setIngredients(e.target.value)} required />
      </div>
      <div>
        <label>Instructions:</label>
        <textarea value={instructions} onChange={(e) => setInstructions(e.target.value)} required></textarea>
      </div>
      <button type="submit">Create Recipe</button>
    </form>
  );
};

export default CreateRecipe;
