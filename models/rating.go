package models

import (
	"time"
)

// Rating represents a user's rating and review for a recipe
type Rating struct {
	ID        string    `json:"id"`
	RecipeID  string    `json:"recipeId"`
	UserID    string    `json:"userId"`
	Score     int       `json:"score"` // 1-5 stars
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// RatingInput represents the data needed to create or update a rating
type RatingInput struct {
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

// NewRating creates a new Rating with the given input and generated ID
func NewRating(id string, recipeID string, userID string, input RatingInput) Rating {
	now := time.Now()
	return Rating{
		ID:        id,
		RecipeID:  recipeID,
		UserID:    userID,
		Score:     input.Score,
		Comment:   input.Comment,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// UpdateRating creates a new Rating with updated fields but preserves the original ID, recipe ID, user ID, and creation time
func UpdateRating(original Rating, input RatingInput) Rating {
	return Rating{
		ID:        original.ID,
		RecipeID:  original.RecipeID,
		UserID:    original.UserID,
		Score:     input.Score,
		Comment:   input.Comment,
		CreatedAt: original.CreatedAt,
		UpdatedAt: time.Now(),
	}
}