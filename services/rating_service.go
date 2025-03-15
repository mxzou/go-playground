package services

import (
	"errors"
	"playground/models"
	"playground/repositories"
)

// RatingService handles business logic for recipe ratings
type RatingService struct {
	repository repositories.RatingRepository
	recipeService *RecipeService
}

// NewRatingService creates a new rating service with the given repositories
func NewRatingService(repository repositories.RatingRepository, recipeService *RecipeService) *RatingService {
	return &RatingService{
		repository:    repository,
		recipeService: recipeService,
	}
}

// GetAllRatings returns all ratings
func (s *RatingService) GetAllRatings() []models.Rating {
	return s.repository.FindAll()
}

// GetRatingByID returns a rating by ID
func (s *RatingService) GetRatingByID(id string) (models.Rating, error) {
	return s.repository.FindByID(id)
}

// GetRatingsByRecipeID returns all ratings for a specific recipe
func (s *RatingService) GetRatingsByRecipeID(recipeID string) []models.Rating {
	return s.repository.FindByRecipeID(recipeID)
}

// GetRatingsByUserID returns all ratings by a specific user
func (s *RatingService) GetRatingsByUserID(userID string) []models.Rating {
	return s.repository.FindByUserID(userID)
}

// CreateRating adds a new rating
func (s *RatingService) CreateRating(recipeID string, userID string, input models.RatingInput) (models.Rating, error) {
	// Verify that the recipe exists
	_, err := s.recipeService.GetRecipeByID(recipeID)
	if err != nil {
		return models.Rating{}, err
	}
	
	// Create the rating
	rating := s.repository.Create(recipeID, userID, input)
	return rating, nil
}

// UpdateRating modifies an existing rating
func (s *RatingService) UpdateRating(id string, userID string, input models.RatingInput) (models.Rating, error) {
	// Verify that the rating exists and belongs to the user
	rating, err := s.repository.FindByID(id)
	if err != nil {
		return models.Rating{}, err
	}
	
	// Check if the rating belongs to the user
	if rating.UserID != userID {
		return models.Rating{}, errors.New("unauthorized: rating belongs to another user")
	}
	
	// Update the rating
	return s.repository.Update(id, input)
}

// DeleteRating removes a rating
func (s *RatingService) DeleteRating(id string, userID string) error {
	// Verify that the rating exists and belongs to the user
	rating, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}
	
	// Check if the rating belongs to the user
	if rating.UserID != userID {
		return errors.New("unauthorized: rating belongs to another user")
	}
	
	// Delete the rating
	return s.repository.Delete(id)
}

// GetAverageRatingForRecipe calculates the average rating score for a recipe
func (s *RatingService) GetAverageRatingForRecipe(recipeID string) float64 {
	ratings := s.repository.FindByRecipeID(recipeID)
	if len(ratings) == 0 {
		return 0
	}
	
	total := 0
	for _, rating := range ratings {
		total += rating.Score
	}
	
	return float64(total) / float64(len(ratings))
}