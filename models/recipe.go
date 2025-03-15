package models

import (
	"time"
)

// Recipe represents a cooking recipe
type Recipe struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Ingredients []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PrepTime    int       `json:"prepTime"` // in minutes
	CookTime    int       `json:"cookTime"` // in minutes
	Servings    int       `json:"servings"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// RecipeInput represents the data needed to create or update a recipe
type RecipeInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Ingredients []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PrepTime    int       `json:"prepTime"`
	CookTime    int       `json:"cookTime"`
	Servings    int       `json:"servings"`
	Tags        []string  `json:"tags"`
}

// NewRecipe creates a new Recipe with the given input and generated ID
func NewRecipe(id string, input RecipeInput) Recipe {
	now := time.Now()
	return Recipe{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		Ingredients: input.Ingredients,
		Instructions: input.Instructions,
		PrepTime:    input.PrepTime,
		CookTime:    input.CookTime,
		Servings:    input.Servings,
		Tags:        input.Tags,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateRecipe creates a new Recipe with updated fields but preserves the original ID and creation time
func UpdateRecipe(original Recipe, input RecipeInput) Recipe {
	return Recipe{
		ID:          original.ID,
		Title:       input.Title,
		Description: input.Description,
		Ingredients: input.Ingredients,
		Instructions: input.Instructions,
		PrepTime:    input.PrepTime,
		CookTime:    input.CookTime,
		Servings:    input.Servings,
		Tags:        input.Tags,
		CreatedAt:   original.CreatedAt,
		UpdatedAt:   time.Now(),
	}
}