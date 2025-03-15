package repositories

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"playground/models"
)

// RatingRepository defines the interface for rating storage operations
type RatingRepository interface {
	FindAll() []models.Rating
	FindByID(id string) (models.Rating, error)
	FindByRecipeID(recipeID string) []models.Rating
	FindByUserID(userID string) []models.Rating
	Create(recipeID string, userID string, input models.RatingInput) models.Rating
	Update(id string, input models.RatingInput) (models.Rating, error)
	Delete(id string) error
}

// InMemoryRatingRepository implements RatingRepository with in-memory storage
type InMemoryRatingRepository struct {
	ratings map[string]models.Rating
	mutex   sync.RWMutex
}

// NewInMemoryRatingRepository creates a new in-memory rating repository
func NewInMemoryRatingRepository() *InMemoryRatingRepository {
	return &InMemoryRatingRepository{
		ratings: make(map[string]models.Rating),
	}
}

// FindAll returns all ratings
func (r *InMemoryRatingRepository) FindAll() []models.Rating {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := make([]models.Rating, 0, len(r.ratings))
	for _, rating := range r.ratings {
		result = append(result, rating)
	}
	return result
}

// FindByID returns a rating by ID
func (r *InMemoryRatingRepository) FindByID(id string) (models.Rating, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	rating, exists := r.ratings[id]
	if !exists {
		return models.Rating{}, errors.New("rating not found")
	}
	return rating, nil
}

// FindByRecipeID returns all ratings for a specific recipe
func (r *InMemoryRatingRepository) FindByRecipeID(recipeID string) []models.Rating {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := make([]models.Rating, 0)
	for _, rating := range r.ratings {
		if rating.RecipeID == recipeID {
			result = append(result, rating)
		}
	}
	return result
}

// FindByUserID returns all ratings by a specific user
func (r *InMemoryRatingRepository) FindByUserID(userID string) []models.Rating {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := make([]models.Rating, 0)
	for _, rating := range r.ratings {
		if rating.UserID == userID {
			result = append(result, rating)
		}
	}
	return result
}

// Create adds a new rating
func (r *InMemoryRatingRepository) Create(recipeID string, userID string, input models.RatingInput) models.Rating {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	id := uuid.New().String()
	rating := models.NewRating(id, recipeID, userID, input)
	
	// Store a copy of the rating (immutable pattern)
	r.ratings[id] = rating
	
	return rating
}

// Update modifies an existing rating
func (r *InMemoryRatingRepository) Update(id string, input models.RatingInput) (models.Rating, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	original, exists := r.ratings[id]
	if !exists {
		return models.Rating{}, errors.New("rating not found")
	}

	// Create a new rating with updated fields (immutable pattern)
	updated := models.UpdateRating(original, input)
	
	// Store the updated rating
	r.ratings[id] = updated
	
	return updated, nil
}

// Delete removes a rating
func (r *InMemoryRatingRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.ratings[id]
	if !exists {
		return errors.New("rating not found")
	}

	delete(r.ratings, id)
	return nil
}