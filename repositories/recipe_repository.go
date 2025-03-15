package repositories

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"playground/models"
)

// RecipeRepository defines the interface for recipe storage operations
type RecipeRepository interface {
	FindAll() []models.Recipe
	FindByID(id string) (models.Recipe, error)
	Create(input models.RecipeInput) models.Recipe
	Update(id string, input models.RecipeInput) (models.Recipe, error)
	Delete(id string) error
}

// InMemoryRecipeRepository implements RecipeRepository with in-memory storage
type InMemoryRecipeRepository struct {
	recipes map[string]models.Recipe
	mutex   sync.RWMutex
}

// NewInMemoryRecipeRepository creates a new in-memory recipe repository
func NewInMemoryRecipeRepository() *InMemoryRecipeRepository {
	return &InMemoryRecipeRepository{
		recipes: make(map[string]models.Recipe),
	}
}

// FindAll returns all recipes
func (r *InMemoryRecipeRepository) FindAll() []models.Recipe {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := make([]models.Recipe, 0, len(r.recipes))
	for _, recipe := range r.recipes {
		result = append(result, recipe)
	}
	return result
}

// FindByID returns a recipe by ID
func (r *InMemoryRecipeRepository) FindByID(id string) (models.Recipe, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	recipe, exists := r.recipes[id]
	if !exists {
		return models.Recipe{}, errors.New("recipe not found")
	}
	return recipe, nil
}

// Create adds a new recipe
func (r *InMemoryRecipeRepository) Create(input models.RecipeInput) models.Recipe {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	id := uuid.New().String()
	recipe := models.NewRecipe(id, input)
	
	// Store a copy of the recipe (immutable pattern)
	r.recipes[id] = recipe
	
	return recipe
}

// Update modifies an existing recipe
func (r *InMemoryRecipeRepository) Update(id string, input models.RecipeInput) (models.Recipe, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	original, exists := r.recipes[id]
	if !exists {
		return models.Recipe{}, errors.New("recipe not found")
	}

	// Create a new recipe with updated fields (immutable pattern)
	updated := models.UpdateRecipe(original, input)
	
	// Store the updated recipe
	r.recipes[id] = updated
	
	return updated, nil
}

// Delete removes a recipe
func (r *InMemoryRecipeRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.recipes[id]
	if !exists {
		return errors.New("recipe not found")
	}

	delete(r.recipes, id)
	return nil
}